package controller

import (
	"fmt"
	"gameComp-Backend/models"
	auth_service "gameComp-Backend/services/auth"
	user_service "gameComp-Backend/services/user"
	"gameComp-Backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req *models.User

	if err := c.BodyParser(&req); err != nil {
		return utils.RespFail(c, err.Error())

	}
	email := req.Email
	password := req.Password

	if email == "" || password == "" {
		return utils.RespFail(c, "Username and Password are required")
	}

	user := models.User{
		Email: email,
	}

	if !auth_service.Login(email, password) {
		fmt.Println(!auth_service.Login(email, password))
		return utils.RespFail(c, "Username or Password is wrong")
	}
	user.FindOne()
	token, _ := utils.GenerateToken(user)

	resp := RegisterSuccessRes{
		Token: token,
	}
	return utils.RespOK(c, resp, "Login Success")

}

func Register(c *fiber.Ctx) error {
	var req *models.User
	if err := c.BodyParser(&req); err != nil {
		utils.RespFail(c, err.Error())
	}
	username := req.Username
	password := req.Password
	email := req.Email
	// school := req.School

	if username == "" || password == "" || email == "" {
		return utils.RespFail(c, "Username, Password, Email, School are required")
	}
	hashedPassword := utils.Encode(password)
	if hashedPassword == "" {
		return utils.RespFail(c, "Password Hash Failed")
	}

	var user = models.User{
		Email: email,
	}

	user.FindOne()

	if user.ID != 0 {
		return utils.RespFail(c, "Email is already in use by another account")
	}

	user = models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	user_service.CreateUser(&user)

	token, _ := utils.GenerateToken(user)

	resp := RegisterSuccessRes{
		Token: token,
	}

	return utils.RespOK(c, resp, "Register Success")

}

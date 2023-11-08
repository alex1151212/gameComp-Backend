package controller

import (
	"encoding/json"
	"fmt"
	"gameComp-Backend/models"
	auth_service "gameComp-Backend/services/auth"
	user_service "gameComp-Backend/services/user"
	"gameComp-Backend/utils"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type googlereCAPTCHARequest struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

type googlereCAPTCHAResponse struct {
	// 確保與JSON回應的屬性相符
	Success bool `json:"success"`
	// 其他回應屬性...
}

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
		return utils.RespUnauthorized(c, "Username or Password is wrong")
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

func GoogleVaild(c *fiber.Ctx) error {
	googlereCAPTCHARequest := new(googlereCAPTCHARequest)

	// 解析請求的JSON主體
	if err := c.BodyParser(googlereCAPTCHARequest); err != nil {
		return err
	}

	googleURL := "https://www.google.com/recaptcha/api/siteverify"
	req, err := http.NewRequest("POST", googleURL, nil)
	if err != nil {
		return err
	}

	// 建立URL查詢參數
	query := req.URL.Query()
	query.Add("secret", googlereCAPTCHARequest.Secret)
	query.Add("response", googlereCAPTCHARequest.Response)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	googlereCAPTCHAResponse := new(googlereCAPTCHAResponse)
	if err := json.Unmarshal(body, &googlereCAPTCHAResponse); err != nil {
		return err
	}

	// 回傳JSON回應
	return c.JSON(googlereCAPTCHAResponse)
}

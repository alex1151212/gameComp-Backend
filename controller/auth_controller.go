package controller

import (
	"encoding/json"
	"gameComp-Backend/models"
	auth_service "gameComp-Backend/services/auth"
	user_service "gameComp-Backend/services/user"
	"gameComp-Backend/utils"
	"io"
	"net/http"
	"net/mail"

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
		return utils.RespFail(c, "Email and Password are required")
	}

	hashedPassword := utils.Md5(password)

	user := auth_service.Login(email, hashedPassword)
	if user == nil {
		return utils.RespUnauthorized(c, "Email or Password is wrong")
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		return utils.RespFail(c, "Generate Token Fail")
	}

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
	password := req.Password
	email := req.Email

	// doing format validate in email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return utils.RespFail(c, "Email format is wrong")
	}

	if password == "" || email == "" {
		return utils.RespFail(c, "Password, Email are required")
	}
	hashedPassword := utils.Md5(password)
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
		Email:    email,
		Password: hashedPassword,
	}

	if err := user_service.CreateUser(&user); err != nil {
		return utils.RespFail(c, "failed to create user")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return utils.RespFail(c, "Generate Token Fail")
	}

	resp := RegisterSuccessRes{
		Token: token,
	}

	return utils.RespOK(c, resp, "Register Success")

}

func GoogleVaild(c *fiber.Ctx) error {
	googlereCAPTCHARequest := new(googlereCAPTCHARequest)

	if err := c.BodyParser(googlereCAPTCHARequest); err != nil {
		return err
	}

	googleURL := "https://www.google.com/recaptcha/api/siteverify"
	req, err := http.NewRequest("POST", googleURL, nil)
	if err != nil {
		return err
	}

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

func ResetPassword(c *fiber.Ctx) error {
	var req *models.User
	if err := c.BodyParser(&req); err != nil {
		utils.RespFail(c, err.Error())
	}
	email := req.Email

	if email == "" {
		return utils.RespFail(c, "Email is required")
	}

	user := models.User{
		Email: email,
	}

	user.FindOne()

	if user.ID == 0 {
		return utils.RespFail(c, "Email is not exist")
	}

	auth_service.ResetPassword(&user)

	return utils.RespOK(c, nil, "Reset Password Success")
}

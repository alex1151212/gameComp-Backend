package middleware

import (
	"encoding/json"
	"fmt"
	"gameComp-Backend/utils"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		return utils.RespUnauthorized(c, "Unauthorized")
	}
	fmt.Println(token)
	user, err := utils.ValidateToken(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		return utils.RespUnauthorized(c, err.Error())
	}

	c.Locals("Auth", user)
	return c.Next()
}

type googlereCAPTCHARequest struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}
type googlereCAPTCHAResponse struct {
	Success bool    `json:"success"`
	Score   float32 `json:"score"`
}

func Recaptcha(c *fiber.Ctx) (err error) {
	secret := c.Get("CaptchaSecret")
	response := c.Get("CaptchaResponse")

	googleURL := "https://www.google.com/recaptcha/api/siteverify"
	req, err := http.NewRequest("POST", googleURL, nil)
	if err != nil {
		return utils.RespImRobot(c)
	}

	query := req.URL.Query()
	query.Add("secret", secret)
	query.Add("response", response)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return utils.RespImRobot(c)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.RespImRobot(c)
	}

	googlereCAPTCHAResponse := new(googlereCAPTCHAResponse)
	if err := json.Unmarshal(body, &googlereCAPTCHAResponse); err != nil {
		return utils.RespImRobot(c)
	}

	if !googlereCAPTCHAResponse.Success || googlereCAPTCHAResponse.Score < 0.5 {
		return utils.RespImRobot(c)
	}
	return c.Next()
}

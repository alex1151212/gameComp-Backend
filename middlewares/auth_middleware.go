package middleware

import (
	"fmt"
	"gameComp-Backend/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		utils.RespUnauthorized(c, "Unauthorized")
	}
	fmt.Println(token)
	user, err := utils.ValidateToken(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		utils.RespUnauthorized(c, err.Error())
	}

	fmt.Println(err)

	c.Locals("Auth", user)
	return c.Next()
}

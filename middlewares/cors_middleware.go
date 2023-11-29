package middleware

import "github.com/gofiber/fiber/v2/middleware/cors"

var CorsMiddleware = cors.New(cors.Config{
	AllowOrigins: "*",
	AllowHeaders: "Origin, Content-Type, Accept, Authorization, CaptchaResponse",
})

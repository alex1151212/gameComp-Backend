package router

import (
	"fmt"
	controller "gameComp-Backend/controller"
	middleware "gameComp-Backend/middlewares"
	"os"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func StartServer() {
	app = fiber.New()
	app.Use(middleware.CorsMiddleware)

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	initUserRoute()
	initAuthRoute()

	fmt.Println("🟢 WebServer start Success.")
	err := app.Listen(host + ":" + port)
	if err != nil {
		return
	}
}

func initAuthRoute() {
	AuthRouter := app.Group("/auth")
	AuthRouter.Use(middleware.AuthMiddleware)

	AuthRouter.Post("/upload", middleware.FileSizeLimitMiddleware(10*1024*1024), controller.UploadGameFile)
	AuthRouter.Get("/users", controller.GetUsers)

}

func initUserRoute() {
	app.Post("/createUser", controller.Register)

	app.Post("/login", controller.Login)
	app.Post("/recaptcha", controller.GoogleVaild)
	app.Get("/users", controller.GetUsers)
	// app.Put("/user/:id", controllers.UpdateUser)
}

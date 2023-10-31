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

	AuthRouter.Post("/upload", controller.UploadGameFile)
	// AuthRouter.Get("", AuthController.Refresh)
	// AuthRouter.Put("/profile", AuthController.EditProfile)
	// AuthRouter.Put("/changePassword", AuthController.ChangePassword)
	// AuthRouter.Get("/servers", ServerController.GetServers)
}

func initUserRoute() {
	// app.Get("/user", controller.GetUsers)
	app.Post("/createUser", controller.Register)

	app.Post("/login", controller.Login)
	app.Get("/users", controller.GetUsers)
	// app.Get("/user", controllers.GetUsers)
	// app.Get("/user/:id", controllers.GetUser)
	// app.Post("/user", controllers.CreateUser)
	// app.Put("/user/:id", controllers.UpdateUser)
	// app.Delete("/user/:id", controllers.DeleteUser)
}

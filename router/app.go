package router

import (
	"fmt"
	controller "gameComp-Backend/controller"
	middleware "gameComp-Backend/middlewares"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var app *fiber.App

func StartServer() {

	app = fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024 * 1024, // this is the default limit of 4MB
	})

	file, err := os.OpenFile("./logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
	}))

	app.Use(middleware.CorsMiddleware)

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	staticPath := os.Getenv("FILE_STORAGE_PATH")
	fmt.Println(staticPath)
	app.Static("/", staticPath)

	initUserRoute()
	initAuthRoute()
	initUtilsRoute()

	fmt.Println("🟢 WebServer start Success.")

	err = app.Listen(host + ":" + port)
	if err != nil {
		return
	}
}

func initAuthRoute() {
	AuthRouter := app.Group("/auth")
	AuthRouter.Use(middleware.AuthMiddleware)
	// middleware.FileSizeLimitMiddleware(10*1024*1024), controller.UploadGameFile)
	AuthRouter.Post("/upload", controller.UploadGameFile)
	AuthRouter.Put("/user", controller.UpdateUser)
	AuthRouter.Post("/team", controller.UserApply)
	AuthRouter.Get("/users", controller.GetUsers)

	AuthRouter.Get("/profile", controller.GetUserProfile)

}

func initUserRoute() {
	app.Post("/createUser", controller.Register)

	app.Post("/login", controller.Login)
	app.Post("/recaptcha", controller.GoogleVaild)
	app.Get("/users", controller.GetUsers)
	// app.Put("/user/:id", controllers.UpdateUser)
}

func initUtilsRoute() {
	UtilsRouter := app.Group("/utils")
	UtilsRouter.Post("/reset/password", controller.ResetPassword)
}

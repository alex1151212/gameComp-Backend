package router

import (
	"fmt"
	controller "gameComp-Backend/controller"
	middleware "gameComp-Backend/middlewares"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var app *fiber.App

func StartServer() {

	app = fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024 * 1024,
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
	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		},
	))

	app.Use(middleware.CorsMiddleware)

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	{
		AuthRouter := app.Group("/auth")
		AuthRouter.Use(middleware.AuthMiddleware)
		AuthRouter.Get("/profile", controller.GetUserProfile)

		AuthRouter.Use(middleware.Recaptcha)
		AuthRouter.Put("/user", controller.UpdateUser)

		AuthRouter.Use(middleware.FileSizeLimitMiddleware(100))
		AuthRouter.Post("/upload", controller.UploadGameFile)
		AuthRouter.Post("/team", controller.UserApply)
	}

	{
		UserRouter := app.Group("")
		UserRouter.Post("/recaptcha", controller.GoogleVaild)

		UserRouter.Use(middleware.Recaptcha)
		UserRouter.Post("/createUser", controller.Register)
		UserRouter.Post("/login", controller.Login)
	}

	{
		staticPath := os.Getenv("FILE_STORAGE_PATH")
		static := app.Group("/")
		static.Use(middleware.AuthMiddleware)
		static.Use(middleware.FilePermissionMiddleware)
		static.Static("/", staticPath, fiber.Static{})
	}

	// initUserRoute()
	// initAuthRoute()
	// initUtilsRoute()

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
	AuthRouter.Get("/profile", controller.GetUserProfile)

	AuthRouter.Use(middleware.Recaptcha)
	AuthRouter.Post("/upload", controller.UploadGameFile)
	AuthRouter.Put("/user", controller.UpdateUser)
	AuthRouter.Post("/team", controller.UserApply)
}

func initUserRoute() {

	app.Post("/recaptcha", controller.GoogleVaild)

	app.Use(middleware.Recaptcha)
	app.Post("/createUser", controller.Register)
	app.Post("/login", controller.Login)
}

func initUtilsRoute() {
	UtilsRouter := app.Group("/utils")
	UtilsRouter.Post("/reset/password", controller.ResetPassword)
}

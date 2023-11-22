package main

import (
	"fmt"
	"gameComp-Backend/db"
	"gameComp-Backend/router"
	"gameComp-Backend/utils"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println(envErr)
	}
	// fmt.Println("\nSwagger Docs: http://127.0.0.1:8888/swagger/index.html")
	db.InitMySQL()
	utils.InitJWTKey()
	utils.AutoMigrate()

	router.StartServer()

}

package auth_service

import (
	"fmt"
	"gameComp-Backend/models"
	"gameComp-Backend/utils"
	"os"
)

func Login(email string, hashedPassword string) bool {
	user := models.User{
		Email: email,
	}
	user.FindOne()
	fmt.Println(user)

	return utils.Compare(user.Password, hashedPassword)
}

func ResetPassword(user *models.User) {
	defaultPassword := os.Getenv("DEFAULT_PASSWORD")
	user.Password = utils.Encode(defaultPassword)
	user.Save()
}

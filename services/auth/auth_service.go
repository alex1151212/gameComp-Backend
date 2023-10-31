package auth_service

import (
	"fmt"
	"gameComp-Backend/models"
	"gameComp-Backend/utils"
)

func Login(email string, hashedPassword string) bool {
	user := models.User{
		Email: email,
	}
	user.FindOne()
	fmt.Println(user)

	return utils.Compare(user.Password, hashedPassword)
}

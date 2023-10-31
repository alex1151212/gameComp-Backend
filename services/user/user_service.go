package user_service

import (
	models "gameComp-Backend/models"
)

func GetUsers() *[]models.User {
	var user models.User

	users := user.FindMany()
	return users
}

func CreateUser(user *models.User) {
	user.Create()
}

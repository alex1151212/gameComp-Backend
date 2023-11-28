package user_service

import (
	models "gameComp-Backend/models"
)

func GetUsers() *[]models.User {
	var user models.User

	users := user.FindMany()
	return users
}

func UpdateUser() *models.User {
	var user models.User

	user.FindOne()
	return &user
}

func CreateUser(user *models.User) error {
	return user.Create().Error
}

func GetUserTeam(user *models.User) models.Team {
	team := models.Team{
		UserID: user.ID,
	}
	team.FindOne()
	return team
}

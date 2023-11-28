package auth_service

import (
	"gameComp-Backend/models"
	"gameComp-Backend/utils"
	"os"
)

func Login(email string, hashedPassword string) *models.User {
	user := models.User{
		Email:    email,
		Password: hashedPassword,
	}
	tx := user.FindOne()
	if tx.Error != nil || tx.RowsAffected != 1 {
		return nil
	}
	return &user
}

func ResetPassword(user *models.User) {
	defaultPassword := os.Getenv("DEFAULT_PASSWORD")
	user.Password = utils.Md5(defaultPassword)
	user.Save()
}

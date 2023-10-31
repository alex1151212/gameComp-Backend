package utils

import (
	"gameComp-Backend/db"
	models "gameComp-Backend/models"
)

func AutoMigrate() {
	db.Instance.AutoMigrate(&models.User{})
}

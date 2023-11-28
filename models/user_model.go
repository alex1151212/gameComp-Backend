package models

import (
	"gameComp-Backend/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(255);not null;unique;" json:"email"`

	Phone    string `gorm:"type:varchar(255);not null;" json:"phone"`
	Password string `gorm:"type:varchar(255);not null;" json:"password"`
	Team     Team   `gorm:"foreignKey:UserID;" json:"team"`
}

func (model *User) Save() *gorm.DB {
	return db.Instance.Save(model)
}

func (model *User) Create() *gorm.DB {
	return db.Instance.Create(model)
}

func (model *User) Delete() *gorm.DB {
	return db.Instance.Delete(&model)
}

func (model *User) FindOne() *gorm.DB {
	return db.Instance.Where(model).First(&model)
}
func (model *User) Update() *gorm.DB {
	return db.Instance.Updates(&model)
}

func (model *User) FindMany() *[]User {
	var user []User
	db.Instance.Find(&user, model)
	return &user
}

package models

import (
	"gameComp-Backend/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255) NOT NULL; unique;" json:"email"`
	Username string `gorm:"type:varchar(255) NOT NULL;" json:"username"`

	Phone    string `gorm:"type:varchar(255) NOT NULL;" json:"phone"`
	Password string `gorm:"type:varchar(255) NOT NULL;" json:"password"`
	IsUpload bool   `gorm:"type:bool NOT NULL;" json:"isUpload"`

	VideoLink string `gorm:"type:varchar(255) NOT NULL;" json:"videoLink"`
	PdfPath   string `gorm:"type:varchar(255) NOT NULL ;" json:"pdfPath"`
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
	return db.Instance.First(&model, &model)
}
func (model *User) Update() *gorm.DB {
	return db.Instance.Updates(&model)
}

func (model *User) FindMany() *[]User {
	var user []User
	db.Instance.Find(&user, model)
	return &user
}

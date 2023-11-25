package models

import (
	"gameComp-Backend/db"

	"gorm.io/gorm"
)

type Attach struct {
	gorm.Model
	UserID       uint   `gorm:"index"`
	Type         string `gorm:"type:varchar(255) NOT NULL;" json:"type"`
	Path         string `gorm:"type:varchar(255) NOT NULL;" json:"path"`
	SaveLocation string `gorm:"type:varchar(255) NOT NULL;" json:"saveLocation"`
	Filename     string `gorm:"type:varchar(255) NOT NULL;" json:"filename"`
}

func (model *Attach) Save() *gorm.DB {
	return db.Instance.Save(model)
}

func (model *Attach) Create() *gorm.DB {
	return db.Instance.Create(model)
}

func (model *Attach) Delete() *gorm.DB {
	return db.Instance.Delete(&model)
}

func (model *Attach) FindOne() *gorm.DB {
	return db.Instance.Where(model).First(&model)
}
func (model *Attach) FindMany() *[]Attach {
	var attach []Attach
	db.Instance.Find(&attach, model)
	return &attach
}

func (model *Attach) Update() *gorm.DB {
	return db.Instance.Updates(&model)
}

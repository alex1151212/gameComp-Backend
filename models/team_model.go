package models

import (
	"gameComp-Backend/db"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	UserID                uint          `gorm:";unique " json:"userID"`
	TeamName              string        `gorm:"type:varchar(255);not null;unique ;" json:"teamName"`
	TeamMember            []TeamMember  `gorm:"serializer:json" json:"teamMember"`
	TeamTeacher           []TeamTeacher `gorm:"serializer:json" json:"teamTeacher"`
	TeamSchoolCertificate []Attach      `gorm:"foreignKey:UserID ;" json:"teamSchoolCertificate"`
	IsUpload              bool          `gorm:"type:bool;not null;default:false;" json:"isUpload"`
	IsApplyTeam           bool          `gorm:"type:bool;not null;default:false;" json:"isApplyTeam"`

	WorkVideoLink string `gorm:"type:varchar(255);" json:"workVideoLink"`
	WorkPdf       Attach `gorm:"foreignKey:UserID;" json:"workPdf"`
}

type TeamMember struct {
	Name string
}

type TeamTeacher struct {
	Name     string
	JobTitle string
}

func (model *Team) Save() *gorm.DB {
	return db.Instance.Save(model)
}

func (model *Team) Create() *gorm.DB {
	return db.Instance.Create(model)
}

func (model *Team) Update() *gorm.DB {
	return db.Instance.Updates(&model)
}

func (model *Team) Delete() *gorm.DB {
	return db.Instance.Delete(&model)
}

func (model *Team) FindOne() *gorm.DB {
	return db.Instance.Where(model).First(&model)
}

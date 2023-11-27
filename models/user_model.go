package models

import (
	"gameComp-Backend/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(255) NOT NULL;unique ;" json:"email"`

	Phone    string `gorm:"type:varchar(255) NOT NULL;" json:"phone"`
	Password string `gorm:"type:varchar(255) NOT NULL;" json:"password"`

	IsUpload bool `gorm:"type:bool NOT NULL;default:false;" json:"isUpload"`

	TeamName              string        `gorm:"type:varchar(255) NOT NULL;unique ;" json:"teamName"`
	TeamMember            []TeamMember  `gorm:"serializer:json" json:"teamMember"`
	TeamTeacher           []TeamTeacher `gorm:"serializer:json" json:"teamTeacher"`
	TeamSchoolCertificate []Attach      `gorm:"foreignKey:UserID ;" json:"teamSchoolCertificate"`

	IsApplyTeam bool `gorm:"type:bool NOT NULL;default:false;" json:"isApplyTeam"`

	WorkVideoLink string `gorm:"type:varchar(255) NOT NULL;" json:"workVideoLink"`
	WorkPdf       Attach `gorm:"foreignKey:UserID ;" json:"workPdf"`
}

type TeamMember struct {
	Name string
}
type TeamTeacher struct {
	Name     string
	JobTitle string
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

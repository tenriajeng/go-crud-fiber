package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	Names    string `json:"names"`
	Post     []Post `gorm:"Foreignkey:UserID;association_foreignkey:ID;" json:"posts"`
}

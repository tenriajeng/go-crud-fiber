package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Names    string `json:"names"`
}

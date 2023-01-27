package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" validate:"required,email" json:"email" `
	Username string `gorm:"unique" json:"username" `
	Password string `validate:"required" json:"-"`
	Names    string `json:"names" `
}

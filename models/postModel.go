package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title     string         `validate:"required,min=6,max=255" json:"title"`
	Body      string         `validate:"required,min=6" json:"body"`
	UserID    uint           `gorm:"column:user_id" json:"user_id"`
	User      User           `validate:"structonly" json:"users"`
}

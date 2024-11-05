package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint      `gorm:"primarykey"`
	Username string    `json:"username" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	Email    string    `json:"email" validate:"required,email" gorm:"unique;"`
	Role     string    `gorm:"default:user"`
	Articles []Article `gorm:"foreignKey:AuthorID"`
}

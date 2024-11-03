package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"not null;unique" json:"name"`
	Description string  `json:"description"`
	Articles    []Article `gorm:"many2many:article_categories" json:"categories"`
}

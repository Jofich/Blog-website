package models

import "gorm.io/gorm"

// add images
type Article struct {
	gorm.Model
	Title      string     `gorm:"not null" json:"title"`
	Content    string     `gorm:"not null" json:"content"`
	AuthorID   uint       `gorm:"index;not null"`
	Author     string     `gorm:"not null"`
	Categories []Category `gorm:"many2many:article_categories" json:"categories"`
}

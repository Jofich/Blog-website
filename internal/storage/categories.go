package storage

import (
	"errors"

	"github.com/Jofich/Blog-website/internal/models"
)

var (
	Categories       map[string]uint
	ErrCategoryExist = errors.New("this category already exist")
)

func (db Storage) LoadCategories() error {
	Categories = make(map[string]uint)
	var cat []models.Category
	result := db.db.Table(CategoryTable).Find(&cat)
	if result.Error != nil {
		return result.Error
	}
	for _, category := range cat {
		Categories[category.Name] = category.ID
	}

	return nil
}

func (db *Storage) SaveCategory(category models.Category) error {

	_, ok := Categories[category.Name]
	if ok {
		return ErrCategoryExist
	}
	if dbc := db.db.Create(&category); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

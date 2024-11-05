package storage

import (
	"github.com/Jofich/Blog-website/internal/models"
)

var Categories map[string]uint

func (db Storage) LoadCategories() error {
	Categories := make(map[string]uint)
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

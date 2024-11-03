package storage

import (
	"fmt"

	"github.com/Jofich/Blog-website/internal/models"
)

var Categories map[string]uint

func (db Storage) LoadCategories() error {
	Categories := make(map[string]uint)
	var cat []models.Category
	result := db.DB.Table(CategoryTable).Find(&cat)
	if result.Error != nil {
		return result.Error
	}
	for _, category := range cat {
		Categories[category.Name] = category.ID
		fmt.Println(category.Name, category.ID)
	}

	return nil
}

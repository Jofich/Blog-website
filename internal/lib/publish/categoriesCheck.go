package publish

import (
	"fmt"
	"strings"

	"github.com/Jofich/Blog-website/internal/models"
	"github.com/Jofich/Blog-website/internal/storage"
)

func IsCategoryExists(articleCategories *[]models.Category) error {
	var incorCat []string
	for _, cat := range *articleCategories {
		if _, ok := storage.Categories[cat.Name]; !ok {
			incorCat = append(incorCat, cat.Name)
		}
	}
	if len(incorCat) != 0 {
		return fmt.Errorf("this categories dont exist: %v", strings.Join(incorCat, ","))
	}
	return nil
}

package categories

import (
	fibererr "github.com/Jofich/Blog-website/internal/lib/fiberErr"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func AddCategory(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		category := c.FormValue("category", "")
		if category == "" {
			return fibererr.Status(c, fiber.StatusBadRequest, "category cant be nil")
		}
		
		return nil
	}
}

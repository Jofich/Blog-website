package user

import (
	"github.com/Jofich/Blog-website/internal/controllers"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func HomePage(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		return c.SendString("Value: " + c.Params(controllers.PathParamUserValue))
	}
}

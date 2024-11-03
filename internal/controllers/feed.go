package controllers

import (
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Default() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.Redirect("/feed", fiber.StatusMovedPermanently)
	}
}

// вывод статей по интересам. Сортировки по дате добавления, оценке, по тегам
func Feed(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		return nil
	}
}

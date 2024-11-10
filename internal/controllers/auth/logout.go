package auth

import (
	"github.com/Jofich/Blog-website/internal/controllers"
	"github.com/Jofich/Blog-website/internal/lib/web/cookies"
	"github.com/gofiber/fiber/v2"
)

func Logout() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		cookies.Delete(c, cookies.JwtName)
		return c.Redirect(controllers.PathFeed, fiber.StatusMovedPermanently)
	}
}

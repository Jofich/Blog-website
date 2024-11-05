package auth

import (
	"github.com/Jofich/Blog-website/internal/controllers"
	authentication "github.com/Jofich/Blog-website/internal/lib/web/cookies/auth"
	"github.com/gofiber/fiber/v2"
)

func Logout() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authentication.Logout(c)
		return c.Redirect(controllers.PathFeed, fiber.StatusMovedPermanently)
	}
}

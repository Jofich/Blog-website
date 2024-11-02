package auth

import (
	authentication "github.com/Jofich/Blog-website/internal/lib/cookies/auth"
	"github.com/gofiber/fiber/v2"
)

func Logout() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authentication.Logout(c)
		return c.Redirect("/home", fiber.StatusMovedPermanently)
	}
}

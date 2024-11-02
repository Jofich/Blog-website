package fibererr

import (
	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx, status int, err string) error {
	if err != "" {
		return c.Status(status).JSON(fiber.Map{
			"error": err,
		})
	}
	c.Status(status)
	return nil
}

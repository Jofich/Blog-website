package fibererr

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx, status int, args ...string) error {
	if len(args) != 0 {
		return c.Status(status).JSON(fiber.Map{
			"error": strings.Join(args, ","),
		})
	}
	c.Status(status)
	return nil
}

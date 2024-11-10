package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx, Name string) {
	c.Cookie(&fiber.Cookie{
		Name:    Name,
		Value:   "",
		Expires: time.Now().Add(-time.Hour * 24),
	})
}

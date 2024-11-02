package auth

import (
	jwtToken "github.com/Jofich/Blog-website/internal/lib/cookies/jwt"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx){
	c.Cookie(&fiber.Cookie{
		Name:   "token",
		MaxAge: -1,
	})
}

func ValidateJWT(c *fiber.Ctx) error {
	token := c.Cookies("token", "")
	if token == "" {
		return jwtToken.ErrTokenEmpty
	}
	err := jwtToken.Valid(token)
	if err != nil {
		Logout(c)
		return err
	}

	return nil
}

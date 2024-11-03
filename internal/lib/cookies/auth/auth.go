package auth

import (
	jwtToken "github.com/Jofich/Blog-website/internal/lib/cookies/jwt"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:   "token",
		MaxAge: -1,
	})
}

func ValidateJWT(c *fiber.Ctx) (models.User, error) {
	token := c.Cookies("token", "")
	if token == "" {
		return models.User{}, jwtToken.ErrTokenEmpty
	}
	user, err := jwtToken.Valid(token)
	if err != nil {
		Logout(c)
		return models.User{}, err
	}

	return user, nil
}

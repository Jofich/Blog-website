package auth

import (
	"github.com/Jofich/Blog-website/internal/lib/web/cookies"
	jwtToken "github.com/Jofich/Blog-website/internal/lib/web/cookies/jwt"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:   cookies.JwtName,
		MaxAge: -1,
	})
}

func ValidateJWT(c *fiber.Ctx) (models.User, error) {
	token := c.Cookies(cookies.JwtName)
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

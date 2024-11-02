package controllers

import (
	"errors"
	"log"

	"github.com/Jofich/Blog-website/internal/lib/cookies/auth"
	jwtToken "github.com/Jofich/Blog-website/internal/lib/cookies/jwt"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Default() (func (c *fiber.Ctx) error){
	return func(c *fiber.Ctx) error {
		return c.Redirect("/home", fiber.StatusMovedPermanently)
	}
}

func Home(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := auth.ValidateJWT(c)
		if err != nil {
			if errors.Is(err, jwtToken.ErrTokenEmpty) {
				return c.Redirect("/signup", fiber.StatusMovedPermanently)
			}
			log.Println(err)
		}

		return nil
	}
}

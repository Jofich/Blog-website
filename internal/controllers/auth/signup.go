package auth

import (
	"errors"
	"log"
	"time"

	"github.com/Jofich/Blog-website/internal/controllers"
	"github.com/Jofich/Blog-website/internal/lib/validator"
	"github.com/Jofich/Blog-website/internal/lib/web/cookies"
	"github.com/Jofich/Blog-website/internal/lib/web/cookies/auth"
	jwtToken "github.com/Jofich/Blog-website/internal/lib/web/cookies/jwt"
	fibererr "github.com/Jofich/Blog-website/internal/lib/web/fiberErr"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func SignupGet(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		_, err := auth.ValidateJWT(c)
		if err != nil {
			if !errors.Is(err, jwtToken.ErrTokenEmpty) {
				log.Println(err)
			}
		} else {
			return c.Redirect(controllers.PathFeed, fiber.StatusMovedPermanently)
		}
		err = c.SendFile("./web/view/signup/index.html")
		if err != nil {
			log.Println(err)
		}
		return err
	}
}
func Signup(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		user := new(models.User)
		err := c.BodyParser(user)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusBadRequest, ErrDataIncorrect)
		}
		err = validator.IsValidUserDataSignup(*user)
		if err != nil {
			return fibererr.Status(c, fiber.StatusBadRequest, err.Error())
		}

		_, err = db.FindUserByEmail(user.Email)
		if err != nil {
			if !errors.Is(err, storage.ErrRecordNotFound) {
				log.Println(err)
				fibererr.Status(c, fiber.StatusInternalServerError, ErrInternalServer)
			}
		} else {
			fibererr.Status(c, fiber.StatusConflict, ErrEmailAlreadyExists)
		}

		_, err = db.FindUserByUsername(user.Username)
		if err != nil {
			if !errors.Is(err, storage.ErrRecordNotFound) {
				log.Println(err)
				fibererr.Status(c, fiber.StatusInternalServerError, ErrInternalServer)
			}
		} else {
			return fibererr.Status(c, fiber.StatusConflict, ErrUsernameAlreadyExists)
		}

		err = db.SaveUser(*user)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusInternalServerError, ErrInternalServer)
		}

		tokenString, err := jwtToken.Create(*user)
		if err != nil {
			log.Println(err)
			fibererr.Status(c, fiber.StatusInternalServerError, ErrInternalServer)
		}
		c.Cookie(&fiber.Cookie{
			Name:   cookies.JwtName,
			Value:  tokenString,
			MaxAge: int(time.Hour.Seconds()) * 10,
		})

		//public/js/signup/script.js check if responce code is 301 and redirect user to "location"
		return c.Status(fiber.StatusMovedPermanently).JSON(fiber.Map{
			"redirect_url": controllers.PathFeed,
			"status":       "success",
			"user": fiber.Map{
				"id":         user.ID,
				"username":   user.Username,
				"email":      user.Email,
				"created_at": user.CreatedAt,
			},
		})
	}
}

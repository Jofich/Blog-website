package auth

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/Jofich/Blog-website/internal/controllers"
	"github.com/Jofich/Blog-website/internal/lib/validator"
	"github.com/Jofich/Blog-website/internal/lib/web/cookies"
	"github.com/Jofich/Blog-website/internal/lib/web/cookies/auth"
	jwtToken "github.com/Jofich/Blog-website/internal/lib/web/cookies/jwt"
	fibererr "github.com/Jofich/Blog-website/internal/lib/web/fiberErr"
	hash "github.com/Jofich/Blog-website/internal/lib/web/hashPassword"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func LoginGet(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		_, err := auth.ValidateJWT(c)
		if err != nil {
			if !errors.Is(err, jwtToken.ErrTokenEmpty) {
				log.Println(err)
			}
		} else {
			return c.Redirect(controllers.PathFeed, fiber.StatusMovedPermanently)
		}

		err = c.SendFile("./web/view/login/index.html")
		if err != nil {
			log.Println(err)
		}
		return err
	}
}

func Login(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		user := new(models.User)

		err := c.BodyParser(user)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusBadRequest, ErrDataIncorrect)
		}
		// user can login by username or email
		// login/email
		// password
		if strings.Contains(user.Username, "@") {
			user.Email = user.Username
			user.Username = ""
		}

		err = validator.IsValidUserDataLogin(*user)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusBadRequest, err.Error())
		}

		password := user.Password
		err = db.UserExist(user)
		if err != nil {
			log.Println(err)
			if err == storage.ErrRecordNotFound {
				return fibererr.Status(c, fiber.StatusBadRequest, ErrUserNotFound)
			}
			return fibererr.Status(c, fiber.StatusInternalServerError, ErrInternalServer)

		}
		if !hash.CompareHashPassword(user.Password, password) {
			return fibererr.Status(c, fiber.StatusBadRequest, ErrAuthentication)
		}
		db.GetUserArticles(user, -1)
		tokenString, err := jwtToken.Create(*user)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusInternalServerError)
		}
		c.Cookie(&fiber.Cookie{
			Name:   cookies.JwtName,
			Value:  tokenString,
			MaxAge: int(time.Hour.Seconds()) * 10,
		})

		return c.Status(fiber.StatusMovedPermanently).JSON(fiber.Map{
			"redirect_url": controllers.PathFeed,
			"status":       "success",
			"message":      "Login success",
		})

	}
}

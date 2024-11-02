package auth

import (
	"log"
	"strings"
	"time"

	"github.com/Jofich/Blog-website/internal/lib/cookies/auth"
	jwtToken "github.com/Jofich/Blog-website/internal/lib/cookies/jwt"
	fibererr "github.com/Jofich/Blog-website/internal/lib/fiberErr"
	hash "github.com/Jofich/Blog-website/internal/lib/hashPassword"
	"github.com/Jofich/Blog-website/internal/lib/validator"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func LoginGet(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := auth.ValidateJWT(c)
		if err != nil {
			log.Println(err)
		} else {
			return c.Redirect("/home", fiber.StatusMovedPermanently)
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
			return fibererr.Status(c,fiber.StatusBadRequest,err.Error())
		}

		
		err = validator.IsValidUserDataLogin(*user)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusBadRequest, err.Error())
		}
		// user can login by username or email
		// login/email
		// password
		if strings.Contains(user.Username, "@") {
			user.Email = user.Username
			user.Username = ""
		}
		password := user.Password
		err = db.UserExist(user)
		if err != nil {
			log.Println(err)
			if err == storage.ErrRecordNotFound {
				return fibererr.Status(c, fiber.StatusBadRequest, "user not found")
			}
			return fibererr.Status(c, fiber.StatusInternalServerError, "something get wrong, please try again")

		}
		if !hash.CompareHashPassword(user.Password, password) {
			return fibererr.Status(c, fiber.StatusBadRequest, "Please verify your password and account name and try again.")
		}

		tokenString, err := jwtToken.Create(user.Username)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusInternalServerError, "")
		}
		c.Cookie(&fiber.Cookie{
			Name:   "token",
			Value:  tokenString,
			MaxAge: int(time.Hour.Seconds()) * 10,
		})

		return c.Status(fiber.StatusMovedPermanently).JSON(fiber.Map{
			"redirect_url": "/home",
			"status": "success",
			"message" : "Login success",
		})

	}
}

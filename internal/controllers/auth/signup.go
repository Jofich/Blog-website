package auth

import (
	"log"
	"time"

	"github.com/Jofich/Blog-website/internal/lib/cookies/auth"
	jwtToken "github.com/Jofich/Blog-website/internal/lib/cookies/jwt"
	fibererr "github.com/Jofich/Blog-website/internal/lib/fiberErr"
	"github.com/Jofich/Blog-website/internal/lib/validator"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func SignupGet(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		err := auth.ValidateJWT(c)
		if err != nil {
			log.Println(err)
		} else {
			return c.Redirect("/home", fiber.StatusMovedPermanently)
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
			return fibererr.Status(c, fiber.StatusBadRequest, "data are incorrect")
		}
		err = validator.IsValidUserDataSignup(*user)
		if err != nil {
			return fibererr.Status(c, fiber.StatusBadRequest, err.Error())
		}
		var existUsers []models.User
		err = db.DB.Where("email = ?", user.Email).Limit(1).Find(&existUsers).Error
		if err != nil {
			log.Println(err)
		}
		if len(existUsers) != 0 {
			return fibererr.Status(c, fiber.StatusConflict, "user with this mail already exists")
		}

		err = db.DB.Where("username = ?", user.Username).Limit(1).Find(&existUsers).Error
		if err != nil {
			log.Println(err)
		}
		if len(existUsers) != 0 {
			return fibererr.Status(c, fiber.StatusConflict, "this username is already taken")
		}

		err = db.SaveUser(*user)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusInternalServerError, "")
		}
		tokenString, err := jwtToken.Create(user.Username)
		if err != nil {
			log.Println(err)
			fibererr.Status(c, fiber.StatusInternalServerError, "something went wrong. try again.")
		}
		c.Cookie(&fiber.Cookie{
			Name:   "token",
			Value:  tokenString,
			MaxAge: int(time.Hour.Seconds()) * 10,
		})

		//public/js/signup/script.js check if responce code is 301 and redirect user to "location"
		return c.Status(fiber.StatusMovedPermanently).JSON(fiber.Map{
			"redirect_url": "/home",
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

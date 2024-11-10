package admin

import (
	"errors"
	"log"

	"github.com/Jofich/Blog-website/internal/lib/web/cookies/auth"
	jwtToken "github.com/Jofich/Blog-website/internal/lib/web/cookies/jwt"
	fibererr "github.com/Jofich/Blog-website/internal/lib/web/fiberErr"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func AddCategory(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		user, err := auth.ValidateJWT(c)
		if err != nil {
			if errors.Is(err, jwtToken.ErrTokenEmpty) {
				return c.Redirect("/login")
			} else {
				log.Println(err)
				return fibererr.Status(c, fiber.StatusBadRequest, "something get wrong,try again")
			}
		}
		if user.Role != models.RoleAdmin && user.Role != models.RoleSuperAdmin {
			return fibererr.Status(c, fiber.StatusMethodNotAllowed, "you are not authorized to execute this command")
		}

		category := new(models.Category)
		err = c.BodyParser(category)
		if err != nil {
			return fibererr.Status(c, fiber.StatusBadRequest, "can not parse request")
		}

		err = db.SaveCategory(*category)
		if err != nil {
			log.Println(err)
			if errors.Is(err, storage.ErrCategoryExist) {
				return fibererr.Status(c, fiber.StatusBadRequest, storage.ErrCategoryExist.Error())
			}
			return fibererr.Status(c, fiber.StatusInternalServerError, "something get wrong,try again")
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Category created successfully",
		})
	}
}

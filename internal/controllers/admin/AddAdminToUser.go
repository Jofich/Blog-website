package admin

import (
	"errors"
	"log"

	"github.com/Jofich/Blog-website/internal/lib/web/cookies/auth"
	jwtToken "github.com/Jofich/Blog-website/internal/lib/web/cookies/jwt"
	fibererr "github.com/Jofich/Blog-website/internal/lib/web/fiberErr"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/gofiber/fiber/v2"
)

func AddAdminToUser() func(c *fiber.Ctx) error {
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
		if user.Role != models.RoleSuperAdmin {
			return fibererr.Status(c, fiber.StatusMethodNotAllowed, "you are not authorized to execute this command")
		}
		return nil
	}
}

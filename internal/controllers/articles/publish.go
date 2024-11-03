package articles

import (
	"log"

	"github.com/Jofich/Blog-website/internal/lib/cookies/auth"
	fibererr "github.com/Jofich/Blog-website/internal/lib/fiberErr"
	"github.com/Jofich/Blog-website/internal/lib/publish"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Publish(db storage.Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		user, err := auth.ValidateJWT(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "you must be logged in to submit an article",
			})
		}
		article := new(models.Article)
		err = c.BodyParser(article)
		if err != nil {
			log.Println(err)
			return fibererr.Status(c, fiber.StatusBadRequest, "")
		}
		err = publish.IsCategoryExists(&article.Categories)
		if err != nil {
			return fibererr.Status(c, fiber.StatusBadRequest, err.Error())
		}
		article.Author = user.Username
		article.AuthorID = user.ID
		db.SaveArtical(*article)

		return nil
	}
}

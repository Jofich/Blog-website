package routes

import (
	"github.com/Jofich/Blog-website/internal/controllers"
	"github.com/Jofich/Blog-website/internal/controllers/admin"
	"github.com/Jofich/Blog-website/internal/controllers/articles"
	"github.com/Jofich/Blog-website/internal/controllers/auth"
	"github.com/Jofich/Blog-website/internal/controllers/user"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, db storage.Storage) {

	app.Get(controllers.PathDefault, controllers.Default())
	app.Post(controllers.PathLogin, auth.Login(db))
	app.Post(controllers.PathSignup, auth.Signup(db))
	app.Post(controllers.PathPublishArticles, articles.Publish(db))
	app.Get(controllers.PathFeed, controllers.Feed(db))
	app.Get(controllers.PathLogout, auth.Logout())
	app.Get(controllers.PathSignup, auth.SignupGet(db))
	app.Get(controllers.PathLogin, auth.LoginGet(db))
	app.Post(controllers.PathAddCategory, admin.AddCategory(db))
	app.Get(controllers.PathUserPage, user.HomePage(db))

}

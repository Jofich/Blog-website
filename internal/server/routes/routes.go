package routes

import (
	"github.com/Jofich/Blog-website/internal/controllers"
	"github.com/Jofich/Blog-website/internal/controllers/auth"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, db storage.Storage) {

	app.Get("/", controllers.Default())
	app.Post("/login", auth.Login(db))
	app.Post("/signup", auth.Signup(db))
	app.Get("/home", controllers.Home(db))
	app.Get("/logout", auth.Logout())
	app.Get("/signup", auth.SignupGet(db))
	app.Get("/login", auth.LoginGet(db))

}

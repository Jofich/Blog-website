package server

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Jofich/Blog-website/internal/config"
	"github.com/Jofich/Blog-website/internal/server/routes"
	"github.com/Jofich/Blog-website/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start(cfg *config.ServerCfg, storage storage.Storage) {

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	app := fiber.New(fiber.Config{})

	app.Use(logger.New())

	app.Static("/static", "./web/public/")
	routes.AuthRoutes(app, storage)

	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))

	log.Fatalln(app.Listen(addr))

}

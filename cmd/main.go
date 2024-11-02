package main

import (
	"github.com/Jofich/Blog-website/internal/config"
	"github.com/Jofich/Blog-website/internal/server"
	"github.com/Jofich/Blog-website/internal/storage/postgres"
)


func main() {

	cfg := config.Load()
	storage := postgres.InitStorage(cfg.DBCfg)
	server.Start(&cfg.ServerCfg, storage)
}

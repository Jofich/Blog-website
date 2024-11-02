package postgres

import (
	"fmt"
	"log"

	"github.com/Jofich/Blog-website/internal/config"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/Jofich/Blog-website/internal/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitStorage(cfg config.DBCfg) storage.Storage {

	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.Login, cfg.Password, cfg.DB_name, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln(err)
	}

	return storage.Storage{DB: db}
}
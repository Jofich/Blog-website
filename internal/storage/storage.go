package storage

import (
	"fmt"
	"time"

	hash "github.com/Jofich/Blog-website/internal/lib/hashPassword"
	"github.com/Jofich/Blog-website/internal/models"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type Storage struct {
	DB *gorm.DB
}

func (s *Storage) SaveUser(u models.User) error {
	var err error
	user := u
	user.Password, err = hash.GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}
	s.DB.Create(&user)
	return nil
}

func (s *Storage) UserExist(u *models.User) error {

	now := time.Now()
	var err error
	if u.Username != "" {
		err = s.DB.Table("users").Where("username = ?", u.Username).Limit(1).Find(u).Error
	} else {
		err = s.DB.Table("users").Where("email = ?", u.Email).Limit(1).Find(u).Error
	}
	duration := time.Since(now)
	fmt.Println("UserExist: ", duration.Seconds())
	if err != nil {
		return err
	}
	return nil
}

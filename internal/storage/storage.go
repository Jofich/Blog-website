package storage

import (
	"fmt"

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
	err = s.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FindUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	err := s.DB.Table(UserTable).Where("username = ?", username).Limit(1).Find(user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (s *Storage) FindUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := s.DB.Table(UserTable).Where("email = ?", email).Limit(1).Find(user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (s *Storage) UserExist(u *models.User) error {
	var user *models.User
	var err error
	if u.Username != "" {
		user, err = s.FindUserByUsername(u.Username)
	} else {
		user, err = s.FindUserByEmail(u.Email)
	}
	if err != nil {
		return err
	}
	*u = *user
	fmt.Println(*u)

	return nil
}

func (s *Storage) SaveArtical(article models.Article) error {

	err := s.DB.Table(ArticalTable).Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

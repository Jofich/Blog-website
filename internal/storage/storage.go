package storage

import (
	"errors"
	"fmt"

	hash "github.com/Jofich/Blog-website/internal/lib/web/hashPassword"
	"github.com/Jofich/Blog-website/internal/models"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type Storage struct {
	db *gorm.DB
}

func Init(stor *gorm.DB) Storage {
	return Storage{db: stor}
}

func (s *Storage) SaveUser(u models.User) error {
	var err error
	user := u
	user.Password, err = hash.GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}
	err = s.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FindUserByUsername(username string) (*models.User, error) {
	var users []models.User
	err := s.db.Table(UserTable).Where("username = ?", username).Limit(1).Find(&users).Error
	if err != nil {
		return &models.User{}, err
	}
	if len(users) == 0 {
		return &models.User{}, ErrRecordNotFound
	}
	return &users[0], nil
}

func (s *Storage) FindUserByEmail(email string) (*models.User, error) {
	var users []models.User
	err := s.db.Table(UserTable).Where("email = ?", email).Limit(1).Find(&users).Error
	if err != nil {
		return &models.User{}, err
	}
	if len(users) == 0 {
		return &models.User{}, ErrRecordNotFound
	}
	return &users[0], nil
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
	return nil
}

func (s *Storage) SaveArtical(article models.Article) error {

	err := s.db.Table(ArticalTable).Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

// 0 or -1 will get all user Articles
func (s *Storage) GetUserArticles(user *models.User, limit int) error {
	var Articles []models.Article
	if limit < -1 {
		return errors.New("limit cant be less than -1")
	}
	if limit == 0 {
		limit = -1
	}
	err := s.db.Table(ArticalTable).Where("author_id = ?", user.ID).Limit(limit).Find(&Articles).Error
	if err != nil {
		return err
	}
	fmt.Println(Articles)
	return nil
}

package models

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string
	Role     string
	ID       uint
}

func (c Claims) Valid() error {
	if c.Username == "" || c.Role == "" || c.ID == 0 {
		return errors.New("")
	}
	return nil
}

func (c *Claims) FromUser(user User) {
	*c = Claims{
		Username: user.Username,
		ID:       user.ID,
		Role:     user.Role,
	}
}
func (c Claims) ToUser() User{
	return  User{
		Username: c.Username, 
		ID: c.ID,
		Role: c.Role,
	}
}
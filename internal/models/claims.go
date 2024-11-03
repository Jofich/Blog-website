package models

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string
	ID       uint
}

func (c Claims) Valid() error {
	if c.Username == "" {
		return errors.New("")
	}
	return nil
}

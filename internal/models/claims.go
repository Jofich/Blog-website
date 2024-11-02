package models

import (
	"errors"
)

type Claims struct {
	Username string
}

func (c Claims) Valid() error {
	if c.Username == "" {
		return errors.New("")
	}
	return nil
}

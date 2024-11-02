package jwtToken

import (
	"errors"

	"github.com/Jofich/Blog-website/config"
	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrTokenEmpty   = errors.New("token is empty")
)

// returns username parsed from jwt token
func Valid(token string) error {
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})
	if err != nil {
		return err
	}
	if _, ok := tokenParsed.Claims.(jwt.MapClaims); ok && tokenParsed.Valid {
		return nil
	} else {
		return ErrInvalidToken
	}

}

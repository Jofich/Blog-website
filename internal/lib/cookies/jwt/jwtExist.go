package jwtToken

import (
	"errors"

	"github.com/Jofich/Blog-website/config"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrTokenEmpty   = errors.New("token is empty")
)

// returns username parsed from jwt token
func Valid(token string) (models.User, error) {
	tokenParsed, err := jwt.ParseWithClaims(token, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})
	if err != nil {
		return models.User{}, err
	}
	if claims, ok := tokenParsed.Claims.(*models.Claims); ok && tokenParsed.Valid {
		return models.User{ID: claims.ID, Username: claims.Username}, nil
	} else {
		return models.User{}, ErrInvalidToken
	}

}

package jwtToken

import (
	"github.com/Jofich/Blog-website/config"
	"github.com/Jofich/Blog-website/internal/models"
	"github.com/golang-jwt/jwt"
)

func Create(username string,id uint) (string, error) {
	claims := models.Claims{
		Username: username,
		ID: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

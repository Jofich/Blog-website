package hash

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	now := time.Now()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	duration := time.Since(now)
	fmt.Println("GeneratePas: ", duration.Seconds())
	return string(bytes), err
}

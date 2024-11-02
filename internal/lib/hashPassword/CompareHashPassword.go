package hash

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CompareHashPassword(hash, password string) bool {
	now := time.Now()



	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))


	duration := time.Since(now)
	fmt.Println("CompareHashPassword: ", duration.Seconds())
	return err == nil
}

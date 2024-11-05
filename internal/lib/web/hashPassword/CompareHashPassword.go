package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashPassword(hash, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

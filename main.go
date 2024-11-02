package main

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrInvalidCharacters = errors.New("contains invalid characters")

func validateString(input string) error {
	re := regexp.MustCompile(`^[a-zA-Z\d][a-zA-Z\d_-]*$`)
	if !re.MatchString(input) {
		return ErrInvalidCharacters
	}
	return nil
}

func main() {
	if err := validateString("1234567890BCDE FGHIJKLMNOPQRSTUVWXYZab_cdefghijklmnopqrstuvwxyz"); err != nil {
		if errors.Is(err, ErrInvalidCharacters) {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Строка допустима")
	}
}

package validator

import (
	"errors"
	"regexp"

	"github.com/Jofich/Blog-website/internal/models"
	valid "github.com/go-playground/validator/v10"
)

const (
	MinUserNameLen = 4
	MaxUserNameLen = 20
	MinPasswordLen = 6
	MaxPasswrodLen = 32
)

var (
	ErrEmailUsernameMissing = errors.New("username or email missing")
	ErrUsernameMissing      = errors.New("username missing")
	ErrPasswordMissing      = errors.New("password missing")
	ErrEmailMissing         = errors.New("email missing")

	ErrEmailInvalid    = errors.New("email invalid")
	ErrDataInvalid     = errors.New("data  contains invalid characters")
	ErrPasswordInvalid = errors.New("password  contains invalid characters")

	ErrUsernameTooShort = errors.New("username is too short")
	ErrUsernameTooLong  = errors.New("username is too long")
	ErrPasswordTooShort = errors.New("password is too short")
	ErrPasswordTooLong  = errors.New("password is too long")

	ErrUsernameInvalid = errors.New("username contains invalid characters")
)

func isEmailValid(email string) error {
	u := models.User{Email: email}
	validate := valid.New()
	err := validate.Struct(u)
	if err != nil {
		return ErrEmailInvalid
	}
	return nil
}

func isUsernameValid(username string) error {

	lenUserName := len(username)

	if lenUserName < MinUserNameLen {
		return ErrUsernameTooShort
	}
	if lenUserName > MaxUserNameLen {
		return ErrUsernameTooLong
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`)
	if !re.MatchString(username) {
		return ErrUsernameInvalid
	}
	return nil
}

func isPasswordValid(password string) error {
	if len(password) > MaxPasswrodLen {
		return ErrPasswordTooLong
	}
	if len(password) < MinPasswordLen {
		return ErrPasswordTooShort
	}
	re := regexp.MustCompile(`^[\w\S]+$`)
	if !re.MatchString(password) {
		return ErrPasswordInvalid
	}
	return nil
}

func IsValidUserDataSignup(user models.User) error {
	if user.Email == "" {
		return ErrEmailMissing
	}
	if user.Username == "" {
		return ErrUsernameMissing
	}
	if user.Password == "" {
		return ErrPasswordMissing
	}
	if err := isEmailValid(user.Email); err != nil {
		return err
	}
	if err := isUsernameValid(user.Username); err != nil {
		return err
	}
	if err := isPasswordValid(user.Password); err != nil {
		return err
	}
	return nil
}

func IsValidUserDataLogin(user models.User) error {
	if user.Username == "" && user.Email != "" {
		if err := isEmailValid(user.Email); err != nil {
			return err
		}

	} else if user.Username != "" && user.Email == "" {
		if err := isUsernameValid(user.Username); err != nil {
			return err
		}
	} else if user.Username == "" && user.Email == "" {
		return ErrEmailUsernameMissing
	} else {
		return ErrDataInvalid
	}
	err := isPasswordValid(user.Password)
	if err != nil {
		return err
	}
	return nil
}

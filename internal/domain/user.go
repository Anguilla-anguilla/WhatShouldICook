package domain

import (
	"regexp"
	"strings"
)

type User struct {
	ID           int64
	UserName     string
	Email        string
	PasswordHash string
}

func (u *User) Validate() (err error) {
	if err = u.validateName(); err != nil {
		return err
	}
	if err = u.validateEmail(); err != nil {
		return err
	}
	if err = u.validatePassword(); err != nil {
		return err
	}
	return
}

func (u *User) validateName() error {
	str := strings.ReplaceAll(u.UserName, " ", "")
	if str == "" {
		return ErrEmptyName
	}
	return nil
}

func (u *User) validateEmail() error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return ErrInvalidEmail
	}
	return nil
}

func (u *User) validatePassword() error {
	if u.PasswordHash == "" {
		return ErrEmptyPassword
	}
	return nil
}

package domain

import "errors"

var (
	ErrNotFound      = errors.New("Error: entity not found")
	ErrAlreadyExists = errors.New("Error: entity already exists")
	ErrUnauthorized  = errors.New("Error: missing permission")
)

var (
	ErrEmptyName        = errors.New("Error: no name provided")
	ErrEmptyDescription = errors.New("Error: no description provided")
	ErrEmptyPassword    = errors.New("Error: no password provided")
	ErrInvalidEmail     = errors.New("Error: invalid email")
	ErrInvalidPassword  = errors.New("Error: invalid password")
	ErrPasswordTooShort = errors.New("Error: password is shorter than 8 symbols")
	ErrInvalidScore     = errors.New("Error: value is less or more than intended")
)

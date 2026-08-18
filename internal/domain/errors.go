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
	ErrEmptyCategory    = errors.New("Error: no category provided")
	ErrEmptyCuisine     = errors.New("Error: no cuisine provided")
	ErrEmptyRation      = errors.New("Error: no ration provided")
	ErrInvalidDuration  = errors.New("Error: invalid duration")
	ErrInvalidEmail     = errors.New("Error: invalid email")
	ErrInvalidUser      = errors.New("Error: invalid user")
	ErrInvalidPassword  = errors.New("Error: invalid password")
	ErrPasswordTooShort = errors.New("Error: password is shorter than 8 symbols")
	ErrInvalidRange     = errors.New("Error: value is out of range")
	ErrNoRecipes        = errors.New("Error: no recipes provided")
)

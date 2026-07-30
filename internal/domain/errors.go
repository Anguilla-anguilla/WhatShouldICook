package domain

import "errors"

var (
	ErrNotFound      = errors.New("Error: entity not found")
	ErrAlreadyExists = errors.New("Error: entity already exists")
	ErrUnauthorized = errors.New("Error: missing permission")
)

var (
	ErrEmptyName        = errors.New("Error: no name provided")
	ErrEmptyDescription = errors.New("Error: no description provided")
)

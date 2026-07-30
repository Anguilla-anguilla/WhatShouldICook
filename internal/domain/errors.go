package domain

import "errors"

var (
	ErrNotFound      = errors.New("Error: entity not found")
	ErrAlreadyExists = errors.New("Error: entity already exists")
)

var (
	ErrEmptyName        = errors.New("Error: no name provided")
	ErrEmptyDescription = errors.New("Error: no description provided")
)

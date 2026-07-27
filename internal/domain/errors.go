package domain

import "errors"

var (
	ErrEmptyName = errors.New("Error: no name provided")
	ErrNotFound = errors.New("Error: entity not found")
)


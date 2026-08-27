package domain

import (
	"fmt"
	"strings"
	"time"
)

type Recipe struct {
	ID              int64
	Name            string
	UserID          int64
	Description     string
	CookingTime     int64
	Price           float64
	ExpiresAfter    int64
	StoreInFreezer  bool
	Favorite        bool
	FridgelessStore int64
	Public          bool
	// Hidden			bool
	CategoryID int64
	CuisineID  int64
	CreatedAt  time.Time
}

func (r *Recipe) Validate() error {
	if r.UserID <= 0 {
		fmt.Printf("In validation %v", r.UserID)
		return ErrInvalidUser
	}
	if err := r.validateName(); err != nil {
		return err
	}
	if err := r.validateFridgelessStore(); err != nil {
		return err
	}
	if err := r.validateFK(r.CategoryID, r.CuisineID); err != nil {
		return err
	}
	return nil
}

func (r *Recipe) validateFK(category, cuisine int64) error {
	if category <= 0 {
		return ErrEmptyCategory
	}
	if cuisine <= 0 {
		return ErrEmptyCuisine
	}
	return nil
}

func (r *Recipe) validateName() error {
	name := strings.ReplaceAll(r.Name, " ", "")
	if name == "" {
		return ErrEmptyName
	}
	return nil
}

func (r *Recipe) validateFridgelessStore() error {
	if r.FridgelessStore < 0 || r.FridgelessStore > 3 {
		return ErrInvalidRange
	}
	return nil
}

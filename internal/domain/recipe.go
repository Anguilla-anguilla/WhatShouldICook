package domain

import (
	"strings"
	"time"
)

// ДОБАВИТЬ НОВЫЕ МИГРАЦИИ
type Recipe struct {
	ID              int64
	Name            string //not null
	UserID          int64
	Description     string // убрать not null. Если блюдо состоит из одного банана - то не нужно ничего писать
	CookingTime     int64
	Price           float64
	ExpiresAfter    int64
	StoreInFreezer  bool
	Favorite        bool
	FridgelessStore int64 // сделаю бальную градацию от 0 до 3, в зависимости от выживаемости еды
	Public          bool
	// Hidden			bool потом добавлю, а то и так много пока всего
	CategoryID int64 // not null
	CuisineID  int64
	CreatedAt  time.Time
}

func (r *Recipe) Validate() error {
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
	if category == 0 {
		return ErrEmptyCategory
	}
	if cuisine == 0 {
		return  ErrEmptyCuisine
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

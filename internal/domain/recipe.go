package domain

import "strings"

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
	CreatedAt  int64
}

func (r *Recipe) Validate() error {
	if err := r.ValidateName(); err != nil {
		return err
	}
	if err := r.ValidateFridgelessStore(); err != nil {
		return err
	}
	return nil
}

func (r *Recipe) ValidateName() error {
	name := strings.ReplaceAll(r.Name, " ", "")
	if name == "" {
		return ErrEmptyName
	}
	return nil
}

func (r *Recipe) ValidateFridgelessStore() error {
	if r.FridgelessStore < 0 || r.FridgelessStore > 3 {
		return ErrInvalidRange
	}
	return nil
}

package domain

import "strings"

// ДОБАВИТЬ НОВЫЕ МИГРАЦИИ
// BOOL can_be_stored_without_fridge
type Recipe struct {
	ID              int64
	Name            string
	IngredientID    int64
	Description     string // мэй би убрать not null?
	CookingTime     int64
	Price           float64
	ExpiresAfter    int64
	StoreInFreezer  bool
	Favorite        bool
	FridgelessStore int64 // сделаю бальную градацию от 0 до 2, в зависимости от выживаемости еды
	CategoryID      int64
	CuisineID       int64
}

// Validate:
// NotNull: ingredient (потом)


// хз, наверное стоит переделать
func (r *Recipe) ValidateNameDescription() error {
	name := strings.ReplaceAll(r.Name, " ", "")
	description := strings.ReplaceAll(r.Description, " ", "")
	if name == "" {
		return ErrEmptyName
	}
	if description == "" {
		return ErrEmptyDescription
	}
	return nil
}

// Обозначить, как enum
// func (r *Recipe) ValidateFridgelessStore() error {
// 	if r.FridgelessStore < 0 || r.FridgelessStore > 3 {
// 		return ErrInvalidScore
// 	}
// }
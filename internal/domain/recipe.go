package domain

type Recipe struct {
	ID   int64
	Name string
	// Ingredient	   Ingredient manytomany
	Description    string
	CookingTime    int64
	Price          float64
	StoreInFreezer bool
	ExpiresAfter   int64
	Favorite       bool
	Category       Category
	Cuisine        Cuisine
}

// Validate:
// NotNull: ingredient (потом)

func (r *Recipe) Validate() error {
	if r.Name == "" {
		return ErrEmptyName
	}
	if r.Description == "" {
		return ErrEmptyDescription
	}
	return nil
}

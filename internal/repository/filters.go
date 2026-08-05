package repository

type RecipeFilters struct {
	UserID     int64
	CategoryID *int64
	CuisineID  *int64
	Favorite   *bool
	Public     *bool
	// Hidden	   *bool
}

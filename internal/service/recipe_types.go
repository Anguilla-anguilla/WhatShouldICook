package service

type CreateRecipeRequest struct {
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	CookingTime     int64               `json:"cooking_time"`
	Price           float64             `json:"price"`
	ExpiresAfter    int64               `json:"expires_after"`
	StoreInFreezer  bool                `json:"store_in_freezer"`
	Favorite        bool                `json:"favorite"`
	FridgelessStore int64               `json:"fridgeless_store"`
	Public          bool                `json:"public"`
	CategoryID      int64               `json:"category_id"`
	CuisineID       int64               `json:"cuisine_id"`
	Ingredients     []IngredientRequest `json:"ingredients"`
}

// Потом добавить юнит
type IngredientRequest struct {
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
}

type RecipeResponse struct {
	ID int64 `json:"id"`
	CreateRecipeRequest
	IngredientsRes []IngredientResponse `json:"ingredients_resp"`
}

type IngredientResponse struct {
	ID int64 `json:"id"`
	IngredientRequest
}

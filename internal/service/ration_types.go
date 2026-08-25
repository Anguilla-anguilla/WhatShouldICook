package service

type RationResponse struct {
	ID       int64              `json:"id"`
	Duration int64              `json:"duration"`
	Recipes  []RecipesForRation `json:"recipes"`
}

type RecipesForRation struct {
	ID   int64 `json:"recipe_id"`
	Name string `json:"name"`
}

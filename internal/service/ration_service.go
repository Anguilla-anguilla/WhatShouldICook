package service

import (
	"context"
)

type RationService struct {
	repo          RationRepository
	repoRR        RationRecipeRepository
	recipeService *RecipeService
}

func NewRationService(repo RationRepository, repoRR RationRecipeRepository, recipeService *RecipeService) *RationService {
	return &RationService{
		repo:          repo,
		repoRR:        repoRR,
		recipeService: recipeService}
}

func (s *RationService) GetByID(ctx context.Context, id, userID int64) (*RationResponse, error) {
	ration, err := s.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	rr, err := s.repoRR.ListByRation(ctx, id)
	if err != nil {
		return nil, err
	}

	recipes := make([]RecipesForRation, 0, len(rr))

	for _, recipeID := range rr {
		recipe, err := s.recipeService.GetByID(ctx, recipeID, userID)
		if err != nil {
			return nil, err
		}
		recipeRes := RecipesForRation{
			ID:   recipeID,
			Name: recipe.Name,
		}
		recipes = append(recipes, recipeRes)
	}
	responce := RationResponse{
		ID:       id,
		Duration: ration.Duration,
		Recipes:  recipes,
	}
	return &responce, nil
}

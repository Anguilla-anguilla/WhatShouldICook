package service

import (
	"WhatShouldICook/internal/domain"
	"context"
	"math/rand"
)

type MenuService struct {
	recipeRepo       RecipeRepository
	rationRepo       RationRepository
	rationRecipeRepo RationRecipeRepository
	shoppingListRepo ShoppingListRepository
	ingredientRepo   IngredientRepository
}

// func (s *MenuService) GenerateMenu(ctx context.Context, req GenerateMenuRequest) (*MenuResponse, error) {

// }

func (s *MenuService) selectRecipes(ctx context.Context, req GenerateMenuRequest) ([]*domain.Recipe, error) {
	var menu []*domain.Recipe

	filters := RecipeFilters{}
	filters.UserID = req.UserID
	filters.CuisineID = &req.CuisineID

	for key, value := range req.CategoryCount {
		catID := key
		filters.CategoryID = &catID

		list, err := s.recipeRepo.List(ctx, filters)
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			continue
		}

		for i := value; i > 0; i-- {
			idx := rand.Intn(len(list))
			menu = append(menu, list[idx])
		}
	}
	return menu, nil
}

func (s *MenuService) createRation(ctx context.Context, userID, duration int64) (*domain.Ration, error) {
	ration := domain.Ration {
		UserID: userID,
		Duration: duration,
	}

	if err := s.rationRepo.Create(ctx, &ration); err != nil {
		return nil, err
	}
	
	return &ration, nil
}

func (s *MenuService) addRecipesToRation(ctx context.Context, rationID int64, recipes []*domain.Recipe) error {
	for _, recipe := range recipes {
		if err := s.rationRecipeRepo.Add(ctx, rationID, recipe.ID); err != nil {
			return err
		}
	}
	return nil
}

func (s *MenuService) buildShoppingList(ctx context.Context, recipes []*domain.Recipe) (domain.ShoppingList, error) {
	
}

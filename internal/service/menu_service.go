package service

import (
	"WhatShouldICook/internal/domain"
	"context"
	"math/rand"
)

type MenuService struct {
	recipeRepo             RecipeRepository
	recipeIngredientRepo   RecipeIngredientRepository
	rationRepo             RationRepository
	rationRecipeRepo       RationRecipeRepository
	shoppingListRepo       ShoppingListRepository
	ingredientRepo         IngredientRepository
	shoppingListRecipeRepo ShoppingListRecipeRepository
}

func NewMenuService(
	recipeRepo RecipeRepository,
	recipeIngredientRepo RecipeIngredientRepository,
	rationRepo RationRepository,
	rationRecipeRepo RationRecipeRepository,
	shoppingListRepo ShoppingListRepository,
	shoppingListRecipeRepo ShoppingListRecipeRepository,
	ingredientRepo IngredientRepository,
) *MenuService {
	return &MenuService{
		recipeRepo:             recipeRepo,
		recipeIngredientRepo:   recipeIngredientRepo,
		rationRepo:             rationRepo,
		rationRecipeRepo:       rationRecipeRepo,
		shoppingListRepo:       shoppingListRepo,
		shoppingListRecipeRepo: shoppingListRecipeRepo,
		ingredientRepo:         ingredientRepo,
	}
}

func (s *MenuService) GenerateMenu(ctx context.Context, req GenerateMenuRequest) (*MenuResponse, error) {
	recipes, err := s.selectRecipes(ctx, req)
	if err != nil {
		return nil, err
	}

	ration, err := s.createRation(ctx, req.UserID, req.Duration)
	if err != nil {
		return nil, err
	}

	err = s.addRecipesToRation(ctx, ration.ID, recipes)
	if err != nil {
		return nil, err
	}

	err = s.buildShoppingList(ctx, recipes, ration.ID)
	if err != nil {
		return nil, err
	}

	shoppingItems, err := s.buildShoppingItems(ctx, recipes)
	if err != nil {
		return nil, err
	}

	var responce MenuResponse
	responce.Ration = ration
	responce.Recipes = recipes
	responce.ShoppingList = shoppingItems

	return &responce, nil
}

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
	ration := domain.Ration{
		UserID:   userID,
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

// Мэй би вообще его удалить и все, что с ним связано
func (s *MenuService) buildShoppingList(ctx context.Context, recipes []*domain.Recipe, rationID int64) error {
	if len(recipes) == 0 {
		return domain.ErrNoRecipes
	}

	var shoppingList domain.ShoppingList
	shoppingList.RationID = rationID
	if err := s.shoppingListRepo.Create(ctx, &shoppingList); err != nil {
		return err
	}

	for _, recipe := range recipes {
		if err := s.shoppingListRecipeRepo.Add(ctx, shoppingList.ID, recipe.ID); err != nil {
			return err
		}
	}

	return nil
}

// ДОБАВИТЬ ЮНИТЫ (А то пользователю будет неудобно с граммовкой)
func (s *MenuService) buildShoppingItems(ctx context.Context, recipes []*domain.Recipe) ([]ShoppingItem, error) {
	var shoppingItems []ShoppingItem
	count := make(map[string]int64)

	for _, recipe := range recipes {
		ingredientIDs, err := s.recipeIngredientRepo.ListByRecipe(ctx, recipe.ID)
		if err != nil {
			return nil, err
		}

		for _, ingredientID := range ingredientIDs {
			ingredient, err := s.ingredientRepo.GetByID(ctx, ingredientID.IngredientID)
			if err != nil {
				return nil, err
			}

			_, ok := count[ingredient.Name]
			if !ok {
				count[ingredient.Name] = ingredientID.Quantity
			} else {
				count[ingredient.Name] += ingredientID.Quantity
			}
		}

	}

	for key, value := range count {
		var shoppingItem ShoppingItem
		shoppingItem.IngredientName = key
		shoppingItem.TotalQuantity = value
		shoppingItems = append(shoppingItems, shoppingItem)
	}
	return shoppingItems, nil
}

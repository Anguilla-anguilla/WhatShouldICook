package service

import (
	"WhatShouldICook/internal/domain"
	"context"
)

type RecipeService struct {
	repo              RecipeRepository
	repoRI            RecipeIngredientRepository
	ingredientService *IngredientService
	cuisineService    *CuisineService
	categoryService   *CategoryService
}

func NewRecipeService(
	repo RecipeRepository,
	repoRI RecipeIngredientRepository,
	ingredientService *IngredientService,
	cuisineService *CuisineService,
	categoryService *CategoryService,
) *RecipeService {
	return &RecipeService{
		repo:              repo,
		repoRI:            repoRI,
		ingredientService: ingredientService,
		cuisineService:    cuisineService,
		categoryService:   categoryService,
	}
}

func (s *RecipeService) Create(ctx context.Context, recipeReq CreateRecipeRequest, userID int64) (*domain.Recipe, error) {

	recipe := &domain.Recipe{
		UserID:          userID,
		Name:            recipeReq.Name,
		Description:     recipeReq.Description,
		CookingTime:     recipeReq.CookingTime,
		Price:           recipeReq.Price,
		ExpiresAfter:    recipeReq.ExpiresAfter,
		StoreInFreezer:  recipeReq.StoreInFreezer,
		Favorite:        recipeReq.Favorite,
		FridgelessStore: recipeReq.FridgelessStore,
		Public:          recipeReq.Public,
		CategoryID:      recipeReq.CategoryID,
		CuisineID:       recipeReq.CuisineID,
	}

	if err := recipe.Validate(); err != nil {
		return nil, err
	}
	if found, _ := s.repo.GetByName(ctx, recipe.Name, userID); found != nil {
		return nil, domain.ErrAlreadyExists
	}

	if _, err := s.cuisineService.GetByID(ctx, recipe.CuisineID, userID); err != nil {
		return nil, err
	}

	if _, err := s.categoryService.GetByID(ctx, recipe.CategoryID); err != nil {
		return nil, err
	}
	if len(recipeReq.Ingredients) == 0 {
		return nil, domain.ErrNoIngredients
	}

	if err := s.repo.Create(ctx, recipe); err != nil {
		return nil, err
	}

	for _, ingredient := range recipeReq.Ingredients {
		ingr, err := s.ingredientService.GetOrCreate(ctx, ingredient.Name)
		if err != nil {
			return nil, err
		}
		if err := s.repoRI.Add(ctx, recipe.ID, ingr.ID, ingredient.Quantity); err != nil {
			return nil, err
		}
	}
	return recipe, nil
}

func (s *RecipeService) List(ctx context.Context, filters RecipeFilters) ([]*RecipeResponce, error) {
	recipes, err := s.repo.List(ctx, filters)
	if err != nil {
		return nil, err
	}

	result := make([]*RecipeResponce, 0, len(recipes))
	for _, recipe := range recipes {
		ri, err := s.repoRI.ListByRecipe(ctx, recipe.ID)
		if err != nil {
			return nil, err
		}

		ingredients := make([]IngredientResponce, 0, len(ri))
		for _, ingID := range ri {
			ingredient, err := s.ingredientService.repo.GetByID(ctx, ingID.IngredientID)
			if err != nil {
				return nil, err
			}
			responce := IngredientResponce{
				ID: ingredient.ID,
				IngredientRequest: IngredientRequest{
					Name:     ingredient.Name,
					Quantity: ingID.Quantity,
				},
			}
			ingredients = append(ingredients, responce)
		}

		recipeRes := RecipeResponce{
			ID: recipe.ID,
			CreateRecipeRequest: CreateRecipeRequest{
				Name:            recipe.Name,
				Description:     recipe.Description,
				CookingTime:     recipe.CookingTime,
				Price:           recipe.Price,
				ExpiresAfter:    recipe.ExpiresAfter,
				StoreInFreezer:  recipe.StoreInFreezer,
				Favorite:        recipe.Favorite,
				FridgelessStore: recipe.FridgelessStore,
				Public:          recipe.Public,
				CategoryID:      recipe.CategoryID,
				CuisineID:       recipe.CuisineID,
			},
			IngredientsRes: ingredients,
		}
		result = append(result, &recipeRes)
	}
	return result, nil
}

func (s *RecipeService) GetByID(ctx context.Context, id, userID int64) (*RecipeResponce, error) {
	recipe, err := s.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	ri, err := s.repoRI.ListByRecipe(ctx, recipe.ID)
	if err != nil {
		return nil, err
	}

	ingredients := make([]IngredientResponce, 0, len(ri))
	for _, ingID := range ri {
		ingredient, err := s.ingredientService.repo.GetByID(ctx, ingID.IngredientID)
		if err != nil {
			return nil, err
		}
		responce := IngredientResponce{
			ID: ingredient.ID,
			IngredientRequest: IngredientRequest{
				Name:     ingredient.Name,
				Quantity: ingID.Quantity,
			},
		}
		ingredients = append(ingredients, responce)
	}
	recipeRes := RecipeResponce{
		ID: recipe.ID,
		CreateRecipeRequest: CreateRecipeRequest{
			Name:            recipe.Name,
			Description:     recipe.Description,
			CookingTime:     recipe.CookingTime,
			Price:           recipe.Price,
			ExpiresAfter:    recipe.ExpiresAfter,
			StoreInFreezer:  recipe.StoreInFreezer,
			Favorite:        recipe.Favorite,
			FridgelessStore: recipe.FridgelessStore,
			Public:          recipe.Public,
			CategoryID:      recipe.CategoryID,
			CuisineID:       recipe.CuisineID,
		},
		IngredientsRes: ingredients,
	}
	return &recipeRes, nil
}

func (s *RecipeService) Update(ctx context.Context, recipeReq CreateRecipeRequest, id, userID int64) error {

	recipe := &domain.Recipe{
		ID:              id,
		Name:            recipeReq.Name,
		UserID:          userID,
		Description:     recipeReq.Description,
		CookingTime:     recipeReq.CookingTime,
		Price:           recipeReq.Price,
		ExpiresAfter:    recipeReq.ExpiresAfter,
		StoreInFreezer:  recipeReq.StoreInFreezer,
		Favorite:        recipeReq.Favorite,
		FridgelessStore: recipeReq.FridgelessStore,
		Public:          recipeReq.Public,
		CategoryID:      recipeReq.CategoryID,
		CuisineID:       recipeReq.CuisineID,
	}

	if err := recipe.Validate(); err != nil {
		return err
	}
	if found, _ := s.repo.GetByName(ctx, recipe.Name, userID); found != nil && found.ID != id {
		return domain.ErrNotFound
	}

	if len(recipeReq.Ingredients) == 0 {
		return domain.ErrNoIngredients
	}

	if err := s.repo.Update(ctx, recipe); err != nil {
		return err
	}

	if err := s.repoRI.DeleteByRecipe(ctx, recipe.ID); err != nil {
		return err
	}
	for _, ingredient := range recipeReq.Ingredients {
		ingr, err := s.ingredientService.GetOrCreate(ctx, ingredient.Name)
		if err != nil {
			return err
		}
		if err := s.repoRI.Add(ctx, recipe.ID, ingr.ID, ingredient.Quantity); err != nil {
			return err
		}
	}

	return nil
}

func (s *RecipeService) Delete(ctx context.Context, id, userID int64) error {
	err := s.repoRI.DeleteByRecipe(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id, userID)
}

func (s *RecipeService) Copy(ctx context.Context, id, userID, ownerID, cuisineID int64) (*domain.Recipe, error) {
	recipe, err := s.GetByID(ctx, id, ownerID)
	if err != nil {
		return nil, err
	}

	if !recipe.Public {
		return nil, domain.ErrPermissionDenied
	}

	newRecipe := &domain.Recipe{
		Name:            recipe.Name,
		UserID:          userID,
		Description:     recipe.Description,
		CookingTime:     recipe.CookingTime,
		Price:           recipe.Price,
		ExpiresAfter:    recipe.ExpiresAfter,
		StoreInFreezer:  recipe.StoreInFreezer,
		Favorite:        false,
		FridgelessStore: recipe.FridgelessStore,
		Public:          false,
		CategoryID:      recipe.CategoryID,
		CuisineID:       cuisineID,
	}

	if err := s.repo.Create(ctx, newRecipe); err != nil {
		return nil, err
	}

	ingredients, err := s.repoRI.ListByRecipe(ctx, recipe.ID)
	if err != nil {
		return nil, err
	}

	for _, ingr := range ingredients {
		if err := s.repoRI.Add(ctx, newRecipe.ID, ingr.IngredientID, ingr.Quantity); err != nil {
			return nil, err
		}
	}

	return newRecipe, nil
}

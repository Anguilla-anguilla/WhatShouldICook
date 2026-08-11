package service

import (
	"WhatShouldICook/internal/domain"
	"context"
)

type CreateRecipeRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	CookingTime     int64   `json:"cooking_time"`
	Price           float64 `json:"price"`
	ExpiresAfter    int64   `json:"expires_after"`
	StoreInFreezer  bool    `json:"store_in_freezer"`
	Favorite        bool    `json:"favorite"`
	FridgelessStore int64   `json:"fridgeless_store"`
	Public          bool    `json:"public"`
	CategoryID      int64   `json:"category_id"`
	CuisineID       int64   `json:"cuisine_id"`
}

type RecipeService struct {
	repo RecipeRepository
}

func NewRecipeService(repo RecipeRepository) *RecipeService {
	return &RecipeService{repo: repo}
}

func (s *RecipeService) Create(ctx context.Context, recipeReq CreateRecipeRequest, userID int64) (*domain.Recipe, error) {

	recipe := &domain.Recipe{
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
	if err := s.repo.Create(ctx, recipe); err != nil {
		return nil, err
	}
	return recipe, nil
}

func (s *RecipeService) List(ctx context.Context, filters RecipeFilters) ([]*domain.Recipe, error) {
	return s.repo.List(ctx, filters)
}
func (s *RecipeService) GetByID(ctx context.Context, id, userID int64) (*domain.Recipe, error) {
	return s.repo.GetByID(ctx, id, userID)
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
	return s.repo.Update(ctx, recipe)
}

func (s *RecipeService) Delete(ctx context.Context, id, userID int64) error {
	return s.repo.Delete(ctx, id, userID)
}

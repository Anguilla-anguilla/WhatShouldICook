package service

import (
	"WhatShouldICook/internal/domain"
	"context"
)

type IngredientService struct {
	repo IngredientRepository
}

func NewIngredientService(repo IngredientRepository) *IngredientService {
	return &IngredientService{repo: repo}
}

func (s *IngredientService) GetOrCreate(ctx context.Context, name string) (*domain.Ingredient, error) {
	ingredient, err := s.GetByName(ctx, name)
	if err == domain.ErrNotFound {
		ingredient, err = s.Create(ctx, name)
	}
	if err != nil {
		return nil, err
	}
	return ingredient, nil
}

func (s *IngredientService) Create(ctx context.Context, name string) (*domain.Ingredient, error) {
	ingredient := &domain.Ingredient{Name: name}
	if err := ingredient.Validate(); err != nil {
		return nil, err
	}

	if found, _ := s.repo.GetByName(ctx, ingredient.Name); found != nil {
		return nil, domain.ErrAlreadyExists
	}
	if err := s.repo.Create(ctx, ingredient); err != nil {
		return nil, err
	}
	return ingredient, nil
}

func (s *IngredientService) List(ctx context.Context) ([]*domain.Ingredient, error) {
	return s.repo.List(ctx)
}

func (s *IngredientService) GetByID(ctx context.Context, id int64) (*domain.Ingredient, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *IngredientService) GetByName(ctx context.Context, name string) (*domain.Ingredient, error) {
	return s.repo.GetByName(ctx, name)
}

func (s *IngredientService) Update(ctx context.Context, id int64, name string) error {
	ingredient, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if found, _ := s.repo.GetByName(ctx, name); found != nil && found.ID != id {
		return domain.ErrAlreadyExists
	}

	ingredient.Name = name
	if err = ingredient.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, ingredient)
}

func (s *IngredientService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

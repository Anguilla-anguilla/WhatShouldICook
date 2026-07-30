package service

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/repository/postgres"
	"context"
)

type CuisineService struct {
	repo *postgres.CuisineRepo
}

func NewCuisineService(repo *postgres.CategoryRepo) *CuisineService {
	return &CuisineService{repo: (*postgres.CuisineRepo)(repo)}
}

func (s CuisineService) Create(ctx context.Context, name, description string, userID int64) (*domain.Cuisine, error) {
	cuisine := &domain.Cuisine{
		Name:        name,
		Description: description,
	}

	if err := cuisine.Validate(); err != nil {
		return nil, err
	}

	if found, _ := s.repo.GetByName(ctx, cuisine.Name, userID); found != nil {
		return nil, domain.ErrAlreadyExists
	}
	if err := s.repo.Create(ctx, cuisine); err != nil {
		return nil, err
	}
	return cuisine, nil
}

func (s *CuisineService) GetByID(ctx context.Context, id, userID int64) (*domain.Cuisine, error) {
	return s.repo.GetByID(ctx, id, userID)
}

func (s *CuisineService) List(ctx context.Context, userID int64) ([]*domain.Cuisine, error) {
	return s.repo.List(ctx, userID)
}

func (s *CuisineService) Update(ctx context.Context, cuisine *domain.Cuisine, userID int64) error {
	if err := cuisine.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, cuisine)
}

func (s *CuisineService) Delete(ctx context.Context, id, userID int64) error {
	return s.repo.Delete(ctx, id)
}

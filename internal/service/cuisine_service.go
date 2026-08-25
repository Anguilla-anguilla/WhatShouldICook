package service

import (
	"WhatShouldICook/internal/domain"
	"context"
	"fmt"
)

type CuisineService struct {
	repo CuisineRepository
}

func NewCuisineService(repo CuisineRepository) *CuisineService {
	return &CuisineService{repo: repo}
}

func (s *CuisineService) Create(ctx context.Context, name, description string, userID int64) (*domain.Cuisine, error) {
	cuisine := &domain.Cuisine{
		Name:        name,
		Description: description,
		UserID:      userID,
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
	fmt.Printf("id: %v, user: %v\n", id, userID)
	return s.repo.GetByID(ctx, id, userID)
}

func (s *CuisineService) List(ctx context.Context, userID int64) ([]*domain.Cuisine, error) {
	return s.repo.List(ctx, userID)
}

func (s *CuisineService) Update(ctx context.Context, cuisine *domain.Cuisine, userID int64) error {
	if err := cuisine.Validate(); err != nil {
		return err
	}
	if cuisine.UserID != userID {
		return domain.ErrUnauthorized
	}
	return s.repo.Update(ctx, cuisine)
}

func (s *CuisineService) Delete(ctx context.Context, id, userID int64) error {
	return s.repo.Delete(ctx, id, userID)
}

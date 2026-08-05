package service

import (
	"WhatShouldICook/internal/domain"
	"context"
)

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) List(ctx context.Context) ([]*domain.Category, error) {
	return s.repo.List(ctx)
}

func (s *CategoryService) GetByID(ctx context.Context, id int64) (*domain.Category, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CategoryService) GetByName(ctx context.Context, name string) (*domain.Category, error) {
	return s.repo.GetByName(ctx, name)
}

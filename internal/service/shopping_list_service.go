package service

import (
	"WhatShouldICook/internal/domain"
	"context"
)

type ShoppingListService struct {
	repo ShoppingListRepository
}

func NewShoppingListService(repo ShoppingListRepository) *ShoppingListService {
	return &ShoppingListService{repo: repo}
}

func (s *ShoppingListService) GetByID(ctx context.Context, id, userID int64) (*domain.ShoppingList, error) {
	return s.repo.GetByID(ctx, id, userID)
}

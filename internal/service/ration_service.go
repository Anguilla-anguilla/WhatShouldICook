package service

import (
	"WhatShouldICook/internal/domain"
	"context"
)

type RationService struct {
	repo RationRepository
}

func NewRationService(repo RationRepository) *RationService {
	return &RationService{repo: repo}
}

func (s *RationService) GetByID(ctx context.Context, id, userID int64) (*domain.Ration, error) {
	return s.repo.GetByID(ctx, id, userID)
}

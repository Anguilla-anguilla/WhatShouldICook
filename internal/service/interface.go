package service

import (
	"WhatShouldICook/internal/domain"
	"context"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]*domain.Category, error)
	GetByID(ctx context.Context, id int64) (*domain.Category, error)
	GetByName(ctx context.Context, name string) (*domain.Category, error)
}

type CuisineRepository interface {
	Create(ctx context.Context, cuisine *domain.Cuisine) error
	List(ctx context.Context, userID int64) ([]*domain.Cuisine, error)
	GetByID(ctx context.Context, id, userID int64) (*domain.Cuisine, error)
	GetByName(ctx context.Context, name string, userID int64) (*domain.Cuisine, error)
	Update(ctx context.Context, cuisine *domain.Cuisine) error
	Delete(ctx context.Context, id, userID int64) error
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, userID int64) (*domain.User, error)
	GetByUserName(ctx context.Context, name string) (*domain.User, error)
	GetByEmail(ctx context.Context, name string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int64) error
}

type IngredientRepository interface {
	Create(ctx context.Context, ingredient *domain.Ingredient) error
	List(ctx context.Context) ([]*domain.Ingredient, error)
	GetByID(ctx context.Context, id int64) (*domain.Ingredient, error)
	GetByName(ctx context.Context, name string) (*domain.Ingredient, error)
	Update(ctx context.Context, ingredient *domain.Ingredient) error
	Delete(ctx context.Context, id int64) error
}

type RecipeRepository interface {
	Create(ctx context.Context, recipe *domain.Recipe) error
	List(ctx context.Context, filters RecipeFilters) ([]*domain.Recipe, error)
	GetByID(ctx context.Context, id, userID int64) (*domain.Recipe, error)
	GetByName(ctx context.Context, name string, userID int64) (*domain.Recipe, error)
	Update(ctx context.Context, recipe *domain.Recipe) error
	Delete(ctx context.Context, id, userID int64) error
}

type RationRepository interface{}

// и для дугих тоже

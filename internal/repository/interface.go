package main

import "context"

type RecipeCategoryRepository interface {
	Create(ctx context.Context, recipe *domain.Recipe) error
	Read(ctx context.Context, limit, offset int) ([]*domain.Recipe, error)
	GetByID(ctx context.Context, id int64) (*domain.Recipe)
	Update(ctx context.Context, recipe *domain.Recipe) error
	Delete(ctx context.Context, recipe *domain.Recipe) error
}

type FoodTypeRepository interface {
	Read(ctx context.Context) ([] *domain.FoodType)
}

type RecipeRepository interface {
	Create(ctx context.Context, recipe *domain.Recipe) error
	Read(ctx context.Context, limit, offset int) ([]*domain.Recipe, error)
	GetByID(ctx context.Context, id int64) (*domain.Recipe, error)
	Update(ctx context.Context, recipe *domain.Recipe) error
	Delete(ctx context.Context, recipe *domain.Recipe) error
}

type RationRepository interface {}
// и для дугих тоже

package main

import "context"

type CategoryRepository interface {
	List(ctx context.Context) ([]*domain.Category, error)
	GetByID(ctx context.Context, id int64) (*domain.Category, error)
}


type CuisineRepository interface {
	Create(ctx context.Context, cuisine *domain.Cuisine) error
	List(ctx context.Context, userID int64) ([]*domain.Cuisine, error)
	GetByID(ctx context.Context, id, userID int64) (*domain.Cuisine, error)
	Update(ctx context.Context, cuisine *domain.Cuisine) error
	Delete(ctx context.Context, id int64) error
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

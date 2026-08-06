package postgres

import (
	"WhatShouldICook/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IngredientRepo struct {
	pool *pgxpool.Pool
}

func NewIngredientRepo(pool *pgxpool.Pool) *IngredientRepo {
	return &IngredientRepo{pool: pool}
}

func (r *IngredientRepo) Create(ctx context.Context, ingredient *domain.Ingredient) error {
	query := `
		INSERT INTO ingredient (name)
		VALUES $1
		RETURNING id
		`
	err := r.pool.QueryRow(ctx, query, ingredient.Name).Scan(&ingredient.ID)
	return err
}

func (r *IngredientRepo) List(ctx context.Context) ([]*domain.Ingredient, error) {
	query := `
		SELECT id, name
		FROM ingredient
		ORDER BY id
		`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []*domain.Ingredient
	for rows.Next() {
		var ingredient domain.Ingredient
		err := rows.Scan(
			&ingredient.ID,
			&ingredient.Name,
		)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, &ingredient)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (r *IngredientRepo) GetByID(ctx context.Context, id int64) (*domain.Ingredient, error) {
	query := `
		SELECT id, name
		FROM ingredient
		WHERE id = $1
		`
	var ingredient domain.Ingredient
	err := r.pool.QueryRow(ctx, query, id).Scan(&ingredient.ID, &ingredient.Name)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *IngredientRepo) GetByName(ctx context.Context, name string) (*domain.Ingredient, error) {
	query := `
		SELECT id, name
		FROM ingredient
		WHERE name = $1
		`
	var ingredient domain.Ingredient
	err := r.pool.QueryRow(ctx, query, name).Scan(&ingredient.ID, &ingredient.Name)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *IngredientRepo) Update(ctx context.Context, ingredient *domain.Ingredient) error {
	query := `
	UPDATE ingredient
	SET name $1
	WHERE id = $2
	`
	result, err := r.pool.Exec(ctx, query, ingredient.Name, ingredient.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *IngredientRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM ingredient WHERE id = $1`
	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}
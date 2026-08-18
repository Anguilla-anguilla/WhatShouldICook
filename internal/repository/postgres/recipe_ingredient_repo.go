package postgres

import (
	"WhatShouldICook/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RecipeIngredientRepo struct {
	pool *pgxpool.Pool
}

func NewRecipeIngredientRepo(pool *pgxpool.Pool) *RecipeIngredientRepo {
	return &RecipeIngredientRepo{pool: pool}
}

func (r *RecipeIngredientRepo) Add(ctx context.Context, recipeID, ingredientID, quantity int64) error {
	query := `
		INSERT INTO recipe_ingredient (recipe_id, ingredient_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (recipe_id, ingredient_id) DO NOTHING
		`
	_, err := r.pool.Exec(ctx, query, recipeID, ingredientID, quantity)
	return err
}

func (r *RecipeIngredientRepo) ListByRecipe(ctx context.Context, recipeID int64) ([]*domain.RecipeIngredient, error) {
	query := `SELECT ingredient_id, quantity FROM recipe_ingredient WHERE recipe_id = $1`
	rows, err := r.pool.Query(ctx, query, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var recipeIngredients []*domain.RecipeIngredient
	for rows.Next() {
		var recipeIngredient domain.RecipeIngredient
		recipeIngredient.RecipeID = recipeID
		if err := rows.Scan(
			&recipeIngredient.IngredientID,
			&recipeIngredient.Quantity); err != nil {
			return nil, err
		}
		recipeIngredients = append(recipeIngredients, &recipeIngredient)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return recipeIngredients, rows.Err()
}

func (r *RecipeIngredientRepo) DeleteByRecipe(ctx context.Context, recipeID int64) error {
	query := `DELETE FROM recipe_ingredient WHERE recipe_id = $1`
	_, err := r.pool.Exec(ctx, query, recipeID)
	return err
}

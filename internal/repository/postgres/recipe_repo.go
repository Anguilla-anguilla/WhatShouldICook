package postgres

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/service"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecipeRepo struct {
	pool *pgxpool.Pool
}

func NewRecipeRepo(pool *pgxpool.Pool) *RecipeRepo {
	return &RecipeRepo{pool: pool}
}

func (r *RecipeRepo) Create(ctx context.Context, recipe *domain.Recipe) error {
	query := `
		INSERT INTO recipe (name,
							user_id,
							description, 
							cooking_time, 
							price, 
							expires_after, 
							is_store_in_freezer, 
							is_favorite,
							frigeless_store,
							is_public,
							category_id, 
							cuisine_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, created_at
		`
	err := r.pool.QueryRow(ctx, query,
		recipe.Name,
		recipe.UserID,
		recipe.Description,
		recipe.CookingTime,
		recipe.Price,
		recipe.ExpiresAfter,
		recipe.StoreInFreezer,
		recipe.Favorite,
		recipe.FridgelessStore,
		recipe.Public,
		recipe.CategoryID,
		recipe.CuisineID,
	).Scan(&recipe.ID, &recipe.CreatedAt)
	return err
}

func (r *RecipeRepo) List(ctx context.Context, filters service.RecipeFilters) ([]*domain.Recipe, error) {
	query := `
		SELECT id,
				name,
				user_id,
				description, 
				cooking_time, 
				price, 
				expires_after, 
				is_store_in_freezer, 
				is_favorite,
				frigeless_store,
				is_public,
				category_id, 
				cuisine_id
		FROM recipe
		WHERE user_id = $1
	`
	args := []interface{}{filters.UserID}
	argIndex := 2

	if filters.CategoryID != nil {
		query += fmt.Sprintf(" AND category_id = $%d", argIndex)
		args = append(args, *filters.CategoryID)
		argIndex++
	}

	if filters.CuisineID != nil {
		query += fmt.Sprintf(" AND cuisine_id = $%d", argIndex)
		args = append(args, *filters.CuisineID)
		argIndex++
	}

	if filters.Favorite != nil {
		query += fmt.Sprintf(" AND is_favorite = $%d", argIndex)
		args = append(args, *filters.Favorite)
		argIndex++
	}

	if filters.Public != nil {
		query += fmt.Sprintf(" AND is_public = $%d", argIndex)
		args = append(args, *filters.Public)
		argIndex++
	}

	// to be continued

	query += " ORDER BY name"
	// и так же можно создать сортировку
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*domain.Recipe
	for rows.Next() {
		var recipe domain.Recipe
		err := rows.Scan(
			&recipe.ID,
			&recipe.Name,
			&recipe.UserID,
			&recipe.Description,
			&recipe.CookingTime,
			&recipe.Price,
			&recipe.ExpiresAfter,
			&recipe.StoreInFreezer,
			&recipe.Favorite,
			&recipe.FridgelessStore,
			&recipe.Public,
			&recipe.CategoryID,
			&recipe.CuisineID,
		)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, &recipe)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *RecipeRepo) GetByID(ctx context.Context, id, userID int64) (*domain.Recipe, error) {
	query := `
		SELECT id,
				name,
				user_id,
				description,
				cooking_time,
				price,
				expires_after,
				is_store_in_freezer,
				is_favorite,
				frigeless_store,
				is_public,
				category_id,
				cuisine_id
		FROM recipe
		WHERE id = $1 AND user_id = $2
	`
	var recipe domain.Recipe
	err := r.pool.QueryRow(ctx, query, id, userID).Scan(
		&recipe.ID,
		&recipe.Name,
		&recipe.UserID,
		&recipe.Description,
		&recipe.CookingTime,
		&recipe.Price,
		&recipe.ExpiresAfter,
		&recipe.StoreInFreezer,
		&recipe.Favorite,
		&recipe.FridgelessStore,
		&recipe.Public,
		&recipe.CategoryID,
		&recipe.CuisineID,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (r *RecipeRepo) GetByName(ctx context.Context, name string, userID int64) (*domain.Recipe, error) {
	query := `
		SELECT id,
				name,
				user_id,
				description,
				cooking_time,
				price,
				expires_after,
				is_store_in_freezer,
				is_favorite,
				frigeless_store,
				is_public,
				category_id,
				cuisine_id
		FROM recipe
		WHERE name = $1 AND user_id = $2
	`
	var recipe domain.Recipe
	err := r.pool.QueryRow(ctx, query, name, userID).Scan(
		&recipe.ID,
		&recipe.Name,
		&recipe.UserID,
		&recipe.Description,
		&recipe.CookingTime,
		&recipe.Price,
		&recipe.ExpiresAfter,
		&recipe.StoreInFreezer,
		&recipe.Favorite,
		&recipe.FridgelessStore,
		&recipe.Public,
		&recipe.CategoryID,
		&recipe.CuisineID,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (r *RecipeRepo) Update(ctx context.Context, recipe *domain.Recipe) error {
	query := `
		UPDATE recipe
		SET name = $1,
			description = $2,
			cooking_time = $3,
			price = $4,
			expires_after = $5,
			is_store_in_freezer = $6,
			is_favorite = $7,
			frigeless_store =$8,
			is_public = $9,
			category_id = $10,
			cuisine_id = $11
		WHERE id = $12 AND user_id = $13
		`
	result, err := r.pool.Exec(ctx, query,
		recipe.Name,
		recipe.Description,
		recipe.CookingTime,
		recipe.Price,
		recipe.ExpiresAfter,
		recipe.StoreInFreezer,
		recipe.Favorite,
		recipe.FridgelessStore,
		recipe.Public,
		recipe.CategoryID,
		recipe.CuisineID,
		recipe.ID,
		recipe.UserID,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *RecipeRepo) Delete(ctx context.Context, id, userID int64) error {
	query := `DELETE FROM recipe WHERE id = $1 AND user_id = $2`
	result, err := r.pool.Exec(ctx, query, id, userID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

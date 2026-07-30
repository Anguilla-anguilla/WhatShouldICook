package postgres

// import (
// 	"WhatShouldICook/internal/domain"
// 	"context"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// // тут нужно много чего переделать
// // после добавления в бд is_public
// type RecipeRepo struct {
// 	pool *pgxpool.Pool
// }

// func NewRecipeRepo(pool *pgxpool.Pool) *RecipeRepo {
// 	return &RecipeRepo{pool: pool}
// }

// func (r RecipeRepo) Create(ctx context.Context, recipe *domain.Recipe) error {
// 	query := `
// 		INSERT INTO recipe (name, description, cooking_time, price, category_id, cuisine_id)
// 		VALUES ($1, $2, $3, $4, $5, $6)
// 		RETURNING id, created_at
// 		`
// 	err := r.pool.QueryRow(ctx, query,
// 		recipe.Name,
// 		recipe.Description,
// 		recipe.CookingTime,
// 		recipe.Price,
// 		recipe.Category,
// 		recipe.Cuisine,
// 	).Scan(&recipe.ID) // &recipe.CreatedAt (?)
// 	return err
// }

// func (r *RecipeRepo) GetByID(ctx context.Context, id int64) (*domain.Recipe, error) {
// 	query := `
// 		SELECT id, name, description, cooking_time, price,
// 		store_in_freezer, expires_after, favorit,
// 		category_id, cuisine_id, created_at
// 		FROM recipe
// 		WHERE id = $1
// 	`
// 	var recipe domain.Recipe
// 	err := r.pool.QueryRow(ctx, query, id).Scan(
// 		&recipe.ID,
// 		&recipe.Name,
// 		&recipe.Description,
// 		&recipe.CookingTime,
// 		&recipe.Price,
// 		&recipe.StoreInFreezer,
// 		&recipe.ExpiresAfter,
// 		&recipe.Favorite,
// 		&recipe.Category,
// 		&recipe.Cuisine,
// 		// &recipe.CreatedAt,
// 	)
// 	if err == pgx.ErrNoRows {
// 		return nil, domain.ErrNotFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &recipe, nil
// }

// func (r *RecipeRepo) List(ctx context.Context) ([]*domain.Recipe, error) {
// 	query := `
// 	`

// 	rows, err := r.pool.Query(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var recipes []*domain.Recipe
// 	for rows.Next() {
// 		var recipe domain.Recipe
// 		err := rows.Scan(
// 			&recipe.ID,
// 			&recipe.Name,
// 			&recipe.Description,
// 			&recipe.CookingTime,
// 			&recipe.Price,
// 			&recipe.StoreInFreezer,
// 			&recipe.ExpiresAfter,
// 			&recipe.Favorite,
// 			&recipe.Category,
// 			&recipe.Cuisine,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		recipes = append(recipes, &recipe)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return recipes, nil
// }

// func (r *RecipeRepo) Update(ctx context.Context, recipe *domain.Recipe) error {
// 	// query := `
// 	// 	`
	
// 	return nil
// }
// func (r *RecipeRepo) Delete(ctx context.Context, recipe *domqin.Recipe) error {
// 	return nil
// }

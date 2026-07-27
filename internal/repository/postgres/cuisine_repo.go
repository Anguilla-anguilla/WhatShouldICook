package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CuisineRepo struct {
	pool *pgxpool.Pool
}

func NewCuisineRepo(pool *pgxpool.Pool) *CuisineRepo {
	return &CuisineRepo{pool: pool}
}

func (c *CuisineRepo) Create(ctx context.Context, cuisine *domain.Cuisine) error {
	query := `
		INSERT INTO cuisine (name, description)
		VALUES ($1, $2)
		RETURNING id
		`
	err := c.pool.QueryRow(ctx, query,
		cuisine.Name,
		cuisine.Description,
	).Scan(&cuisine.ID)
	return err
}

func (c *CuisineRepo) List(ctx context.Context, userID int64) ([]*domain.Cuisine, error) {
	query := `
			SELECT cuisine.id, cuisine.name, cuisine.description
			FROM app_user
			JOIN ration ON app_user.id = ration.user_id
			JOIN ration_recipe ON ration.id = ration_recipe.ration_id
			JOIN recipe ON recipe.id = ration_recipe.recipe_id
			JOIN cuisine ON cuisine.id = recipe.cuisine_id
			WHERE app_user.id = $1
			ORDER BY cuisine.id
			`
	rows, err := c.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cuisines []*domain.Cuisine
	for rows.Next() {
		var cuisine domain.Cuisine
		err := rows.Scan(
			&cuisine.ID,
			&cuisine.Name,
			&cuisine.Description,
		)
		if err != nil {
			return nil, err
		}
		cuisines = append(cuisines, &cuisine)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cuisines, nil
}

func (c *CuisineRepo) GetByID(ctx context.Context, id, userID int64) (*domain.Cuisine, error) {
	query := `
			SELECT cuisine.id, cuisine.name, cuisine.description
			FROM app_user
			JOIN ration ON app_user.id = ration.user_id
			JOIN ration_recipe ON ration.id = ration_recipe.ration_id
			JOIN recipe ON recipe.id = ration_recipe.recipe_id
			JOIN cuisine ON cuisine.id = recipe.cuisine_id
			WHERE app_user.id = $2 AND cuisine.id = $1
			`
	var cuisine domain.Cuisine
	err := c.pool.QueryRow(ctx, query, id, userID).Scan(
		&cuisine.ID,
		&cuisine.Name,
		&cuisine.Description,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &cuisine, nil
}

func (c *CuisineRepo) Update(ctx context.Context, cuisine *domain.Cuisine) error {
	query := `
		UPDATE cuisine 
		SET name = $1,
			description = $2
		WHERE id = $3
		`
	result, err := c.pool.Exec(ctx, query,
		cuisine.Name,
		cuisine.Description,
		cuisine.ID,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (c *CuisineRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM cuisine WHERE id = $q`
	result, err := c.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

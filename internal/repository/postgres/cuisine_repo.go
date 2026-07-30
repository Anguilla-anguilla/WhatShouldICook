package postgres

import (
	"WhatShouldICook/internal/domain"
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
		INSERT INTO cuisine (name, description, user_id)
		VALUES ($1, $2, $3)
		RETURNING id
		`
	err := c.pool.QueryRow(ctx, query,
		cuisine.Name,
		cuisine.Description,
		cuisine.UserID,
	).Scan(&cuisine.ID)
	return err
}

func (c *CuisineRepo) List(ctx context.Context, userID int64) ([]*domain.Cuisine, error) {
	query := `
			SELECT cuisine.id, cuisine.name, cuisine.description, cuisine.user_id
			FROM cuisine
			WHERE user_id = $1
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
			&cuisine.UserID,
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
			SELECT cuisine.id, cuisine.name, cuisine.description, cuisine.user_id
			FROM cuisine
			WHERE user_id = $2 AND cuisine.id = $1
			`
	var cuisine domain.Cuisine
	err := c.pool.QueryRow(ctx, query, id, userID).Scan(
		&cuisine.ID,
		&cuisine.Name,
		&cuisine.Description,
		&cuisine.UserID,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &cuisine, nil
}

func (c *CuisineRepo) GetByName(ctx context.Context, name string, userID int64) (*domain.Cuisine, error) {
	query := `
		SELECT cuisine.id, cuisine.name, cuisine.description, user_id
		FROM cuisine
		WHERE user_id = $2 AND cuisine.name = $1
		`
	var cuisine domain.Cuisine
	err := c.pool.QueryRow(ctx, query, name, userID).Scan(
		&cuisine.ID,
		&cuisine.Name,
		&cuisine.Description,
		&cuisine.UserID,
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
		WHERE id = $3 AND user_id = $4
		`
	result, err := c.pool.Exec(ctx, query,
		cuisine.Name,
		cuisine.Description,
		cuisine.ID,
		cuisine.UserID,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (c *CuisineRepo) Delete(ctx context.Context, id, userID int64) error {
	query := `DELETE FROM cuisine WHERE id = $1 AND user_id = $2`
	result, err := c.pool.Exec(ctx, query, id, userID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

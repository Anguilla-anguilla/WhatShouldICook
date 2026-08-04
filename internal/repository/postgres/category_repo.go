package postgres

import (
	"WhatShouldICook/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepo struct {
	pool *pgxpool.Pool
}

func NewCategoryRepo(pool *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{pool: pool}
}

func (c *CategoryRepo) List(ctx context.Context) ([]*domain.Category, error) {
	query := `SELECT id, name FROM category ORDER BY id`

	rows, err := c.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryRepo) GetByID(ctx context.Context, id int64) (*domain.Category, error) {
	query := `
		SELECT id, name 
		FROM category 
		WHERE id = $1`

	var category domain.Category
	err := c.pool.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryRepo) GetByName(ctx context.Context, name string) (*domain.Category, error) {
	query := `
		SELECT id, name FROM category WHERE name = $1
		`
	var category domain.Category
	err := c.pool.QueryRow(ctx, query, name).Scan(
		&category.ID,
		&category.Name,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &category, nil
}

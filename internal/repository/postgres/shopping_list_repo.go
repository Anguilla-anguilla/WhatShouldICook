package postgres

import (
	"WhatShouldICook/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ShoppingListRepo struct {
	pool *pgxpool.Pool
}

func NewShoppingListRepo(pool *pgxpool.Pool) *ShoppingListRepo {
	return &ShoppingListRepo{pool: pool}
}

func (c *ShoppingListRepo) Create(ctx context.Context, shopping_list *domain.ShoppingList) error {
	query := `
		INSERT INTO shopping_list (ration_id)
		VALUES ($1)
		RETURNING id
		`
	err := c.pool.QueryRow(ctx, query, shopping_list.RationID).Scan(&shopping_list.ID)
	return err
}

func (c *ShoppingListRepo) GetByID(ctx context.Context, id, userID int64) (*domain.ShoppingList, error) {
	query := `
			SELECT id, ration_id
			FROM shopping_list
			WHERE id = $1
			`
	var shopping_list domain.ShoppingList
	err := c.pool.QueryRow(ctx, query, id).Scan(
		&shopping_list.ID,
		&shopping_list.RationID,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &shopping_list, nil
}

func (c *ShoppingListRepo) Delete(ctx context.Context, id, userID int64) error {
	query := `DELETE FROM shopping_list WHERE id = $1`
	result, err := c.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

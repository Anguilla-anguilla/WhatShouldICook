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

func (r *ShoppingListRepo) Create(ctx context.Context, shopping_list *domain.ShoppingList) error {
	query := `
		INSERT INTO shopping_list (ration_id)
		VALUES ($1)
		RETURNING id
		`
	err := r.pool.QueryRow(ctx, query, shopping_list.RationID).Scan(&shopping_list.ID)
	return err
}

func (r *ShoppingListRepo) GetByID(ctx context.Context, id, userID int64) (*domain.ShoppingList, error) {
	query := `
			SELECT shopping_list.id, shopping_list.ration_id
			FROM shopping_list
			JOIN ration ON shopping_list.ration_id = ration.id 
			WHERE shopping_list.id = $1 AND ration.user_id = $2
			`
	var shopping_list domain.ShoppingList
	err := r.pool.QueryRow(ctx, query, id, userID).Scan(
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

func (r *ShoppingListRepo) Delete(ctx context.Context, id, userID int64) error {
	query := `
		DELETE FROM shopping_list
		WHERE id = $1 AND ration_id IN (
			SELECT id FROM ration WHERE user_id = $2
		)
		`
	result, err := r.pool.Exec(ctx, query, id, userID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

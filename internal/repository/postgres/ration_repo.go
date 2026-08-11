package postgres

import (
	"WhatShouldICook/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RationRepo struct {
	pool *pgxpool.Pool
}

func NewRationRepo(pool *pgxpool.Pool) *RationRepo {
	return &RationRepo{pool: pool}
}

func (r *RationRepo) Create(ctx context.Context, ration *domain.Ration) error {
	query := `
		INSERT INTO ration (user_id, duration)
		VALUES ($1, $2)
		RETURNING id
		`
	err := r.pool.QueryRow(ctx, query,
		ration.UserID,
		ration.Duration,
	).Scan(&ration.ID)
	return err
}

func (r *RationRepo) GetByID(ctx context.Context, id, userID int64) (*domain.Ration, error) {
	query := `
			SELECT id, user_id, duration
			FROM ration
			WHERE user_id = $2 AND id = $1
			`
	var ration domain.Ration
	err := r.pool.QueryRow(ctx, query, id, userID).Scan(
		&ration.ID,
		&ration.UserID,
		&ration.Duration,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &ration, nil
}

func (r *RationRepo) Delete(ctx context.Context, id, userID int64) error {
	query := `DELETE FROM ration WHERE id = $1 AND user_id = $2`
	result, err := r.pool.Exec(ctx, query, id, userID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

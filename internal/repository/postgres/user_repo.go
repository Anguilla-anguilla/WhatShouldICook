package postgres

import (
	"WhatShouldICook/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

func (u *UserRepo) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO app_user (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
		`
	err := u.pool.QueryRow(ctx, query,
		user.UserName,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID)
	return err

}

func (u *UserRepo) GetByID(ctx context.Context, userID int64) (*domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		FROM app_user
		WHERE id = $1
		`
	var user domain.User

	err := u.pool.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.PasswordHash,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) GetByUserName(ctx context.Context, name string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		FROM app_user
		WHERE username = $1
		`
	var user domain.User
	err := u.pool.QueryRow(ctx, query, name).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.PasswordHash,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		FROM app_user
		WHERE email = $1
		`
	var user domain.User
	err := u.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.PasswordHash,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE app_user
		SET username = $1,
			email = $2,
			password_hash = $3
		WHERE id = $4
		`
	result, err := u.pool.Exec(ctx, query,
		user.UserName,
		user.Email,
		user.PasswordHash,
		user.ID,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (u *UserRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM app_user WHERE id = $1`
	result, err := u.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

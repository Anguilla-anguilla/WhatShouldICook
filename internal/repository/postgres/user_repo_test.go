package postgres

import (
	"context"
	"testing"

	"WhatShouldICook/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepoCreate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepo(pool)
	ctx := context.Background()

	app_user := &domain.User{
		UserName:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "hash",
	}

	err := repo.Create(ctx, app_user)
	require.NoError(t, err)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM app_user WHERE username = 'testuser'").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "No user found")
}

func TestUserRepoGetByID(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	app_user, err := repo.GetByID(ctx, 1)
	require.NoError(t, err)
	assert.Equal(t, "testuser", app_user.UserName)
}

func TestUserRepoGetByUserName(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	app_user, err := repo.GetByUserName(ctx, "testuser")
	require.NoError(t, err)
	assert.Equal(t, int64(1), app_user.ID)
}

func TestUserRepoGetByEmail(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	app_user, err := repo.GetByEmail(ctx, "test@example.com")
	require.NoError(t, err)
	assert.Equal(t, int64(1), app_user.ID)
}

func TestUserRepoUpdate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	appUserUPD := &domain.User{
		ID:           1,
		UserName:     "testuser",
		Email:        "mail@example.com",
		PasswordHash: "hash",
	}

	err = repo.Update(ctx, appUserUPD)
	require.NoError(t, err)

	var email string
	pool.QueryRow(ctx, "SELECT email FROM app_user WHERE username = 'testuser'").Scan(&email)
	assert.Equal(t, "mail@example.com", email)
}

func TestUserRepoDelete(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	err = repo.Delete(ctx, 1)
	require.NoError(t, err)

	var count int
	pool.QueryRow(ctx, "SELECT COUNT(*) FROM app_user WHERE id = 1").Scan(&count)
	assert.Equal(t, 0, count)
}

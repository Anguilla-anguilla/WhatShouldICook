package postgres

import (
	"context"
	"testing"

	"WhatShouldICook/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRationRepoCreate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRationRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	ration := &domain.Ration{
		UserID:   1,
		Duration: 1,
	}

	err = repo.Create(ctx, ration)
	require.NoError(t, err)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM ration WHERE id = 1").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "No ration found")
}

func TestRationRepoGetByID(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRationRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO ration (user_id, duration)
		VALUES (1, 1)
	`)
	require.NoError(t, err)

	ration, err := repo.GetByID(ctx, 1, 1)
	require.NoError(t, err)
	assert.Equal(t, int64(1), ration.Duration)
}

func TestRationRepoDelete(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRationRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO ration (user_id, duration)
		VALUES (1, 1)
	`)
	require.NoError(t, err)

	err = repo.Delete(ctx, 1, 1)
	require.NoError(t, err)

	var count int
	pool.QueryRow(ctx, "SELECT COUNT(*) FROM ration WHERE id = 1").Scan(&count)
	assert.Equal(t, 0, count)
}

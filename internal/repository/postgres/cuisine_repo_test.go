package postgres

import (
	"context"
	"testing"

	"WhatShouldICook/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCuisineRepoCreate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCuisineRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	cuisine := &domain.Cuisine{
		UserID:      1,
		Name:        "Test Cuisine",
		Description: "Test Description",
	}

	err = repo.Create(ctx, cuisine)
	require.NoError(t, err)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM cuisine WHERE name = 'Test Cuisine'").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "No cuisine found")
}

func TestCuisineRepoList(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCuisineRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (name,
							user_id,
							description)
		VALUES ('Test Cuisine 1', 1, 'Test Description')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (name,
							user_id,
							description)
		VALUES ('Test Cuisine 2', 1, 'Test Description')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (name,
							user_id,
							description)
		VALUES ('Test Cuisine 3', 1, 'Test Description')
	`)
	require.NoError(t, err)

	cuisines, err := repo.List(ctx, 1)
	require.NoError(t, err)

	for i, cuisine := range cuisines {
		assert.Equal(t, int64(i+1), cuisine.ID)
	}
}

func TestCuisineRepoGetByID(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCuisineRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (name,
							user_id,
							description)
		VALUES ('Test Cuisine', 1, 'Test Description')
	`)
	require.NoError(t, err)

	cuisine, err := repo.GetByID(ctx, 1, 1)
	require.NoError(t, err)
	assert.Equal(t, "Test Cuisine", cuisine.Name)
}

func TestCuisineRepoGetByName(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCuisineRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (name,
							user_id,
							description)
		VALUES ('Test Cuisine', 1, 'Test Description')
	`)
	require.NoError(t, err)

	cuisine, err := repo.GetByName(ctx, "Test Cuisine", 1)
	require.NoError(t, err)
	assert.Equal(t, int64(1), cuisine.ID)
}

func TestCuisineRepoUpdate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCuisineRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (name,
							user_id,
							description)
		VALUES ('Test Cuisine', 1, 'Test Description')
	`)
	require.NoError(t, err)

	cuisineUPD := &domain.Cuisine{
		ID:          1,
		UserID:      1,
		Name:        "Test Cuisine",
		Description: "New Description",
	}

	err = repo.Update(ctx, cuisineUPD)
	require.NoError(t, err)

	var description string
	pool.QueryRow(ctx, "SELECT description FROM cuisine WHERE name = 'Test Cuisine'").Scan(&description)

	assert.Equal(t, "New Description", description)
}

func TestCuisineRepoDelete(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCuisineRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (name,
							user_id,
							description)
		VALUES ('Test Cuisine', 1, 'Test Description')
	`)
	require.NoError(t, err)

	err = repo.Delete(ctx, 1, 1)
	require.NoError(t, err)

	var count int
	pool.QueryRow(ctx, "SELECT COUNT(*) FROM cuisine WHERE id = 1").Scan(&count)
	assert.Equal(t, 0, count)
}

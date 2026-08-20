package postgres

import (
	"context"
	"testing"

	"WhatShouldICook/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShoppingListRepoCreate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewShoppingListRepo(pool)
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

	shoppingList := &domain.ShoppingList{RationID: 1}

	err = repo.Create(ctx, shoppingList)
	require.NoError(t, err)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM shopping_list WHERE ration_id = 1").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "No shopping list found")
}

func TestShoppingListRepoGetByID(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewShoppingListRepo(pool)
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

	_, err = pool.Exec(ctx, `
		INSERT INTO shopping_list (ration_id)
		VALUES (1)
	`)
	require.NoError(t, err)

	shoppingList, err := repo.GetByID(ctx, 1, 1)
	require.NoError(t, err)
	assert.Equal(t, int64(1), shoppingList.RationID)
}

func TestShoppingListRepoDelete(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewShoppingListRepo(pool)
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

	_, err = pool.Exec(ctx, `
		INSERT INTO shopping_list (ration_id)
		VALUES (1)
	`)
	require.NoError(t, err)

	err = repo.Delete(ctx, 1, 1)
	require.NoError(t, err)

	var count int
	pool.QueryRow(ctx, "SELECT COUNT(*) FROM shopping_list WHERE id = 1").Scan(&count)
	assert.Equal(t, 0, count)
}

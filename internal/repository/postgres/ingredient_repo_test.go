package postgres

import (
	"context"
	"testing"

	"WhatShouldICook/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIngredientRepoCreate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewIngredientRepo(pool)
	ctx := context.Background()

	ingredient := &domain.Ingredient{
		Name: "Test Ingredient",
	}

	err := repo.Create(ctx, ingredient)
	require.NoError(t, err)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM ingredient WHERE name = 'Test Ingredient'").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "No ingredient found")
}

func TestIngredientRepoList(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewIngredientRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO ingredient (name)
		VALUES ('Test Ingredient 1')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO ingredient (name)
		VALUES ('Test Ingredient 2')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO ingredient (name)
		VALUES ('Test Ingredient 3')
	`)
	require.NoError(t, err)

	ingredients, err := repo.List(ctx)
	require.NoError(t, err)

	for i, ingredient := range ingredients {
		assert.Equal(t, int64(i+1), ingredient.ID)
	}
}

func TestIngredientRepoGetByID(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewIngredientRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO ingredient (name)
		VALUES ('Test Ingredient')
	`)
	require.NoError(t, err)

	ingredient, err := repo.GetByID(ctx, 1)
	require.NoError(t, err)
	assert.Equal(t, "Test Ingredient", ingredient.Name)
}

func TestIngredientRepoGetByName(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewIngredientRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO ingredient (name)
		VALUES ('Test Ingredient')
	`)
	require.NoError(t, err)

	ingredient, err := repo.GetByName(ctx, "Test Ingredient")
	require.NoError(t, err)
	assert.Equal(t, int64(1), ingredient.ID)
}

func TestIngredientRepoUpdate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewIngredientRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO ingredient (name)
		VALUES ('Test Ingredient')
	`)
	require.NoError(t, err)

	ingredientUPD := &domain.Ingredient{
		ID:   1,
		Name: "Test Ingredient new",
	}

	err = repo.Update(ctx, ingredientUPD)
	require.NoError(t, err)

	var name string
	pool.QueryRow(ctx, "SELECT name FROM ingredient WHERE id = 1").Scan(&name)

	assert.Equal(t, "Test Ingredient new", name)
}

func TestIngredientRepoDelete(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewIngredientRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO ingredient (name)
		VALUES ('Test Ingredient')
	`)
	require.NoError(t, err)

	err = repo.Delete(ctx, 1)
	require.NoError(t, err)

	var count int
	pool.QueryRow(ctx, "SELECT COUNT(*) FROM ingredient WHERE id = 1").Scan(&count)
	assert.Equal(t, 0, count)
}

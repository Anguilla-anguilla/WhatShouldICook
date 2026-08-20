package postgres

import (
	"context"
	"testing"

	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecipeRepoCreate(t *testing.T) {

	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRecipeRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test Cuisine')
	`)
	require.NoError(t, err)

	recipe := &domain.Recipe{
		Name:            "Test Recipe",
		UserID:          1,
		Description:     "Test Description",
		CookingTime:     30,
		Price:           10.99,
		ExpiresAfter:    5,
		StoreInFreezer:  true,
		Favorite:        false,
		FridgelessStore: 2,
		Public:          false,
		CategoryID:      1,
		CuisineID:       1,
	}

	err = repo.Create(ctx, recipe)
	require.NoError(t, err)
	assert.NotZero(t, recipe.ID)
	assert.NotZero(t, recipe.CreatedAt)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM recipe WHERE name = 'Test Recipe'").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "No recipe found")
}

func TestRecipeRepoListAndFilter(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRecipeRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test Cuisine')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO recipe (name,
							user_id,
							description, 
							cooking_time, 
							price, 
							expires_after, 
							store_in_freezer, 
							favorite,
							fridgeless_store,
							is_public,
							category_id, 
							cuisine_id)
		VALUES ('Test Recipe 1', 1, 'Test Description', 30, 10.99,
				5, TRUE, FALSE, 2, FALSE, 1, 1)
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO recipe (name,
							user_id,
							description, 
							cooking_time, 
							price, 
							expires_after, 
							store_in_freezer, 
							favorite,
							fridgeless_store,
							is_public,
							category_id, 
							cuisine_id)
		VALUES ('Test Recipe 2', 1, 'Test Description', 30, 10.99,
				5, TRUE, TRUE, 2, FALSE, 1, 1)
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO recipe (name,
							user_id,
							description, 
							cooking_time, 
							price, 
							expires_after, 
							store_in_freezer, 
							favorite,
							fridgeless_store,
							is_public,
							category_id, 
							cuisine_id)
		VALUES ('Test Recipe 3', 1, 'Test Description', 30, 10.99,
				5, TRUE, TRUE, 2, FALSE, 1, 1)
	`)
	require.NoError(t, err)

	fav := true
	var categoryID int64 = 1
	var cuisineID int64 = 1

	filter := service.RecipeFilters{
		UserID:     1,
		CategoryID: &categoryID,
		CuisineID:  &cuisineID,
		Favorite:   &fav}

	recipes, err := repo.List(ctx, filter)
	require.NoError(t, err)

	for i, recipe := range recipes {
		assert.Equal(t, int64(i+2), recipe.ID)
	}
}

func TestRecipeRepoGetByID(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRecipeRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test Cuisine')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO recipe (name,
							user_id,
							description, 
							cooking_time, 
							price, 
							expires_after, 
							store_in_freezer, 
							favorite,
							fridgeless_store,
							is_public,
							category_id, 
							cuisine_id)
		VALUES ('Test Recipe', 1, 'Test Description', 30, 10.99,
				5, TRUE, FALSE, 2, FALSE, 1, 1)
	`)
	require.NoError(t, err)

	recipe, err := repo.GetByID(ctx, 1, 1)
	require.NoError(t, err)
	assert.Equal(t, "Test Recipe", recipe.Name)
}

func TestRecipeRepoGetByName(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRecipeRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test Cuisine')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO recipe (name,
							user_id,
							description, 
							cooking_time, 
							price, 
							expires_after, 
							store_in_freezer, 
							favorite,
							fridgeless_store,
							is_public,
							category_id, 
							cuisine_id)
		VALUES ('Test Recipe', 1, 'Test Description', 30, 10.99,
				5, TRUE, FALSE, 2, FALSE, 1, 1)
	`)
	require.NoError(t, err)

	recipe, err := repo.GetByName(ctx, "Test Recipe", 1)
	require.NoError(t, err)
	assert.Equal(t, int64(1), recipe.ID)
}

func TestRecipeRepoUpdate(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRecipeRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test Cuisine')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO recipe (name,
							user_id,
							description, 
							cooking_time, 
							price, 
							expires_after, 
							store_in_freezer, 
							favorite,
							fridgeless_store,
							is_public,
							category_id, 
							cuisine_id)
		VALUES ('Test Recipe', 1, 'Test Description', 30, 10.99,
				5, TRUE, FALSE, 2, FALSE, 1, 1)
	`)
	require.NoError(t, err)

	recipeUPD := &domain.Recipe{
		ID:              1,
		Name:            "Test Recipe",
		UserID:          1,
		Description:     "New Description",
		CookingTime:     30,
		Price:           12,
		ExpiresAfter:    5,
		StoreInFreezer:  true,
		Favorite:        false,
		FridgelessStore: 2,
		Public:          false,
		CategoryID:      1,
		CuisineID:       1,
	}

	err = repo.Update(ctx, recipeUPD)
	require.NoError(t, err)

	var description string
	var price int64
	pool.QueryRow(ctx, "SELECT description, price FROM recipe WHERE name = 'Test Recipe'").Scan(&description, &price)

	assert.Equal(t, "New Description", description)
	assert.Equal(t, int64(12), price)
}

func TestRecipeRepoDelete(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRecipeRepo(pool)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		INSERT INTO app_user (id, username, email, password_hash) 
		VALUES (1, 'testuser', 'test@example.com', 'hash')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test Cuisine')
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO recipe (name,
							user_id,
							description, 
							cooking_time, 
							price, 
							expires_after, 
							store_in_freezer, 
							favorite,
							fridgeless_store,
							is_public,
							category_id, 
							cuisine_id)
		VALUES ('Test Recipe', 1, 'Test Description', 30, 10.99,
				5, TRUE, FALSE, 2, FALSE, 1, 1)
	`)

	err = repo.Delete(ctx, 1, 1)
	require.NoError(t, err)

	var count int
	pool.QueryRow(ctx, "SELECT COUNT(*) FROM recipe WHERE id = 1").Scan(&count)
	assert.Equal(t, 0, count)
}

package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRationRecipeRepoAdd(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRationRecipeRepo(pool)
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
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test RationRecipe')
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

	err = repo.Add(ctx, 1, 1)
	require.NoError(t, err)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM ration_recipe WHERE ration_id = 1").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "No rr connection found")
}

func TestRationRecipeRepoList(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRationRecipeRepo(pool)
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
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test RationRecipe')
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
		INSERT INTO ration_recipe (ration_id, recipe_id) 
		VALUES (1, 1)
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO ration_recipe (ration_id, recipe_id) 
		VALUES (1, 2)
	`)
	require.NoError(t, err)

	_, err = repo.ListByRation(ctx, 1)
	require.NoError(t, err)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM ration_recipe WHERE ration_id = 1").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 2, count, "Wrong amount of rr connections")
}

func TestRationRecipeRepoDelete(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRationRecipeRepo(pool)
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
		INSERT INTO cuisine (id, user_id, name) 
		VALUES (1, 1, 'Test RationRecipe')
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

	_, err = pool.Exec(ctx, `
		INSERT INTO ration_recipe (ration_id, recipe_id) 
		VALUES (1, 1)
	`)
	require.NoError(t, err)

	err = repo.DeleteByRation(ctx, 1)
	require.NoError(t, err)

	var count int
	pool.QueryRow(ctx, "SELECT COUNT(*) FROM ration_recipe WHERE ration_id = 1").Scan(&count)
	assert.Equal(t, 0, count)
}

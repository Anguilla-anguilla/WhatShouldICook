package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCategoryRepoList(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCategoryRepo(pool)
	ctx := context.Background()

	categories, err := repo.List(ctx)
	require.NoError(t, err)

	assert.Equal(t, len(categories), 6)
}

func TestCategoryRepoGetByID(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCategoryRepo(pool)
	ctx := context.Background()

	category, err := repo.GetByID(ctx, 1)
	require.NoError(t, err)
	assert.Equal(t, "main", category.Name)
}

func TestCategoryRepoGetByName(t *testing.T) {
	pool, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewCategoryRepo(pool)
	ctx := context.Background()

	category, err := repo.GetByName(ctx, "main")
	require.NoError(t, err)
	assert.Equal(t, int64(1), category.ID)
}

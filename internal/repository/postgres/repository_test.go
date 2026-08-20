package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) (*pgxpool.Pool, func()) {
	dsn := "postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), dsn)
	require.NoError(t, err, "Connection to test BD failed")

	return pool, func() {
		_, _ = pool.Exec(context.Background(), "TRUNCATE TABLE recipe_ingredient, recipe, ingredient, cuisine, app_user RESTART IDENTITY CASCADE")
		pool.Close()
	}
}

func cleanup(pool *pgxpool.Pool) {
	_, _ = pool.Exec(context.Background(),
		"TRUNCATE TABLE recipe_ingredient, recipe, ingredient, cuisine, app_user RESTART IDENTITY CASCADE")
}

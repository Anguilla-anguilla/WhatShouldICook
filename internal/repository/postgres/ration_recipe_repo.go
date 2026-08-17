package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RationRecipeRepo struct {
	pool *pgxpool.Pool
}

func NewRationRecipeRepo(pool *pgxpool.Pool) *RationRecipeRepo {
	return &RationRecipeRepo{pool: pool}
}

func (r *RationRecipeRepo) Add(ctx context.Context, rationID, recipeID int64) error {
	query := `
		INSERT INTO ration_recipe (ration_id, recipe_id)
		VALUES ($1, $2)
		ON CONFLICT (ration_id, recipe_id) DO NOTHING
		`
	_, err := r.pool.Exec(ctx, query, rationID, recipeID)
	return err
}

func (r *RationRecipeRepo) ListByRation(ctx context.Context, rationID int64) ([]int64, error) {
	query := `SELECT recipe_id FROM ration_recipe WHERE ration_id = $1`
	rows, err := r.pool.Query(ctx, query, rationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipeIDs []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		recipeIDs = append(recipeIDs, id)
	}
	return recipeIDs, rows.Err()
}

func (r *RationRecipeRepo) DeleteByRation(ctx context.Context, rationID int64) error {
	query := `DELETE FROM ration_recipe WHERE ration_id = $1`
	_, err := r.pool.Exec(ctx, query, rationID)
	return err
}

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ShoppingListRecipeRepo struct {
	pool *pgxpool.Pool
}

func NewShoppingListRecipeRepo(pool *pgxpool.Pool) *ShoppingListRecipeRepo {
	return &ShoppingListRecipeRepo{pool: pool}
}

func (r *ShoppingListRecipeRepo) Add(ctx context.Context, shoppingListID, recipeID int64) error {
	query := `
		INSERT INTO shopping_list_recipe (shopping_list_id, recipe_id)
		VALUES ($1, $2)
		ON CONFLICT (shopping_list_id, recipe_id) DO NOTHING
		`
	_, err := r.pool.Exec(ctx, query, shoppingListID, recipeID)
	return err
}

func (r *ShoppingListRecipeRepo) ListByShoppingList(ctx context.Context, shoppingListID int64) ([]int64, error) {
	query := `SELECT recipe_id FROM shopping_list_recipe WHERE shopping_list_id = $1`
	rows, err := r.pool.Query(ctx, query, shoppingListID)
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

func (r *ShoppingListRecipeRepo) DeleteByShoppingList(ctx context.Context, shoppingListID int64) error {
	query := `DELETE FROM shopping_list_recipe WHERE shopping_list_id = $1`
	_, err := r.pool.Exec(ctx, query, shoppingListID)
	return err
}

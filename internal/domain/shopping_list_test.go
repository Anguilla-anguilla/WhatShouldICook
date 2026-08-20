package domain

import (
	"testing"
)

func TestShoppingListValidate(t *testing.T) {
	tests := []struct {
		name         string
		shoppingList ShoppingList
		wantErr      bool
		errType      error
	}{
		{
			name: "valid shopping list",
			shoppingList: ShoppingList{
				RationID: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid ration",
			shoppingList: ShoppingList{
				RationID: 0,
			},
			wantErr: true,
			errType: ErrEmptyRation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.shoppingList.Validate()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Validate() expected error, got nil")
				}
				if tt.errType != nil && err != tt.errType {
					t.Errorf("Validate() error = %v, want %v", err, tt.errType)
				}
			} else {
				if err != nil {
					t.Errorf("Validate() unexpected error = %v", err)
				}
			}
		})
	}
}

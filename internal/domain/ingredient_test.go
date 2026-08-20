package domain

import (
	"testing"
)

func TestIngredientValidate(t *testing.T) {
	tests := []struct {
		name    string
		ingredient Ingredient
		wantErr bool
		errType error
	}{
		{
			name: "valid ingredient",
			ingredient: Ingredient{
				Name:        "Name",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			ingredient: Ingredient{
				Name:        "",
			},
			wantErr: true,
			errType: ErrEmptyName,
		},
		{
			name: "name with spaces",
			ingredient: Ingredient{
				Name:        "   ",
			},
			wantErr: true,
			errType: ErrEmptyName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ingredient.Validate()

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
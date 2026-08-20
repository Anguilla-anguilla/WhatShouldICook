package domain

import "testing"

func TestRecipeValidate(t *testing.T) {
	tests := []struct {
		name    string
		recipe  Recipe
		wantErr bool
		errType error
	}{
		{
			name: "valid recipe",
			recipe: Recipe{
				UserID:          1,
				CategoryID:      1,
				CuisineID:       1,
				Name:            "Name",
				FridgelessStore: 2,
			},
			wantErr: false,
		},
		{
			name: "invalid userID",
			recipe: Recipe{
				UserID:          0,
				CategoryID:      1,
				CuisineID:       1,
				Name:            "Name",
				FridgelessStore: 2,
			},
			wantErr: true,
			errType: ErrInvalidUser,
		},
		{
			name: "invalid Category",
			recipe: Recipe{
				UserID:          1,
				CategoryID:      0,
				CuisineID:       1,
				Name:            "Name",
				FridgelessStore: 2,
			},
			wantErr: true,
			errType: ErrEmptyCategory,
		},
		{
			name: "invalid Cuisine",
			recipe: Recipe{
				UserID:          1,
				CategoryID:      1,
				CuisineID:       0,
				Name:            "Name",
				FridgelessStore: 2,
			},
			wantErr: true,
			errType: ErrEmptyCuisine,
		},
		{
			name: "empty Name",
			recipe: Recipe{
				UserID:          1,
				CategoryID:      1,
				CuisineID:       0,
				Name:            "   ",
				FridgelessStore: 2,
			},
			wantErr: true,
			errType: ErrEmptyName,
		},
		{
			name: "store too small",
			recipe: Recipe{
				UserID:          1,
				CategoryID:      1,
				CuisineID:       0,
				Name:            "Name",
				FridgelessStore: -1,
			},
			wantErr: true,
			errType: ErrInvalidRange,
		},
		{
			name: "store too big",
			recipe: Recipe{
				UserID:          1,
				CategoryID:      1,
				CuisineID:       0,
				Name:            "Name",
				FridgelessStore: 100,
			},
			wantErr: true,
			errType: ErrInvalidRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.recipe.Validate()

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

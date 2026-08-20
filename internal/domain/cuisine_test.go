package domain

import (
	"testing"
)

func TestCuisineValidate(t *testing.T) {
	tests := []struct {
		name    string
		cuisine Cuisine
		wantErr bool
		errType error
	}{
		{
			name: "valid cuisine",
			cuisine: Cuisine{
				UserID:      1,
				Name:        "Name",
				Description: "Description",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			cuisine: Cuisine{
				UserID:      1,
				Name:        "",
				Description: "Description",
			},
			wantErr: true,
			errType: ErrEmptyName,
		},
		{
			name: "name with spaces",
			cuisine: Cuisine{
				UserID:      1,
				Name:        "   ",
				Description: "Description",
			},
			wantErr: true,
			errType: ErrEmptyName,
		},
		{
			name: "invalid id",
			cuisine: Cuisine{
				UserID: 0,
				Name:   "Name",
			},
			wantErr: true,
			errType: ErrInvalidUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cuisine.Validate()

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

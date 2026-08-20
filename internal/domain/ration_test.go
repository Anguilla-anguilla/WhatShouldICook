package domain

import (
	"testing"
)

func TestRationValidate(t *testing.T) {
	tests := []struct {
		name    string
		ration  Ration
		wantErr bool
		errType error
	}{
		{
			name: "valid ration",
			ration: Ration{
				UserID:   1,
				Duration: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid User ID",
			ration: Ration{
				UserID:   0,
				Duration: 1,
			},
			wantErr: true,
			errType: ErrInvalidUser,
		},
		{
			name: "invalid Duration",
			ration: Ration{
				UserID:   1,
				Duration: -1,
			},
			wantErr: true,
			errType: ErrInvalidDuration,
		},
		{
			name: "invalid Duration and user id",
			ration: Ration{
				UserID:   -1,
				Duration: -0,
			},
			wantErr: true,
			errType: ErrInvalidUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ration.Validate()

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

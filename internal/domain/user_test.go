package domain

import "testing"

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
		errType error
	}{
		{
			name: "valid user",
			user: User{
				UserName:     "Name",
				Email:        "email@email.com",
				PasswordHash: "00000",
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			user: User{
				UserName:     "   ",
				Email:        "email@email.com",
				PasswordHash: "00000",
			},
			wantErr: true,
			errType: ErrEmptyName,
		},
		{
			name: "invalid email 1",
			user: User{
				UserName:     "Name",
				Email:        "email@email.",
				PasswordHash: "00000",
			},
			wantErr: true,
			errType: ErrInvalidEmail,
		},
		{
			name: "invalid email 2",
			user: User{
				UserName:     "Name",
				Email:        "email@.com",
				PasswordHash: "00000",
			},
			wantErr: true,
			errType: ErrInvalidEmail,
		},
		{
			name: "invalid email 3",
			user: User{
				UserName:     "Name",
				Email:        "@email.com",
				PasswordHash: "00000",
			},
			wantErr: true,
			errType: ErrInvalidEmail,
		},
		{
			name: "invalid email 4",
			user: User{
				UserName:     "Name",
				Email:        "email.com",
				PasswordHash: "00000",
			},
			wantErr: true,
			errType: ErrInvalidEmail,
		},
		{
			name: "empty password",
			user: User{
				UserName:     "Name",
				Email:        "email@email.com",
				PasswordHash: "",
			},
			wantErr: true,
			errType: ErrEmptyPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()

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

package service

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}

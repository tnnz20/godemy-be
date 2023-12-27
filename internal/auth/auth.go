package auth

import "context"

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	Token string `json:"access_token"`
}

type Service interface {
	SignIn(ctx context.Context, req *AuthRequest) (*AuthResponse, error)
}

package users

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/users/entities"
)

// TODO: Define Update user data
// Service is a contract that defines the user service.
type Service interface {
	Register(ctx context.Context, req entities.RegisterPayload) (err error)
	Login(ctx context.Context, req entities.LoginPayload) (res entities.LoginResponse, err error)
	GetUser(ctx context.Context, req entities.GetUserPayload) (res entities.UserResponse, err error)
}

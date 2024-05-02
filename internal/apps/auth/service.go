package auth

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/auth/entities"
)

type Service interface {
	Register(ctx context.Context, req entities.RegisterPayload) (err error)
	Login(ctx context.Context, req entities.LoginPayload) (res entities.LoginResponse, err error)
}
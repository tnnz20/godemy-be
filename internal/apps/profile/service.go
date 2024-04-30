package profile

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/profile/entities"
)

type Service interface {
	GetProfileByUserId(ctx context.Context, req entities.GetProfileByUserIdRequest) (res *entities.GetProfileByUserIdResponse, err error)
	// UpdateProfile(ctx context.Context, req entities.UpdateProfileRequest) (err error)
}

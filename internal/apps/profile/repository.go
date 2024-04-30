package profile

import (
	"context"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/profile/entities"
)

type Repository interface {
	FindProfileByUserId(ctx context.Context, userId uuid.UUID) (profile entities.Profile, err error)
	// UpdateProfile(ctx context.Context, profile entities.Profile) (err error)
}

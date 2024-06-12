package users

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/users/entities"
	"github.com/tnnz20/godemy-be/internal/storage/postgres"
)

// TODO: Define Update user data
// Repository is a contract that defines the user repository.
type Repository interface {
	CreateUsersWithTX(ctx context.Context, tx *sql.Tx, user entities.Users) (id uuid.UUID, err error)
	InsertUsersRoleWithTX(ctx context.Context, tx *sql.Tx, role entities.Roles) (err error)
	FindUserByEmail(ctx context.Context, email string) (user entities.Users, err error)
	FindRoleByUserID(ctx context.Context, userID uuid.UUID) (role entities.Roles, err error)
	FindUserByUserId(ctx context.Context, userId uuid.UUID) (user entities.Users, err error)
	UpdateUserProfile(ctx context.Context, users entities.Users) (err error)
	postgres.DBTX
}

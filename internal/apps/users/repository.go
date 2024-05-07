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
	GetUserByEmail(ctx context.Context, email string) (user entities.Users, err error)
	GetRoleByUserID(ctx context.Context, userID uuid.UUID) (role entities.Roles, err error)
	postgres.DBTX
}

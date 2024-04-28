package auth

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/auth/entities"
)

type Repository interface {
	CreateUserWithTX(ctx context.Context, tx *sql.Tx, user entities.User) (id uuid.UUID, err error)
	CreateProfileWithTX(ctx context.Context, tx *sql.Tx, profile entities.Profile) (err error)
	InsertUserRoleWithTX(ctx context.Context, tx *sql.Tx, user entities.User) (err error)
	GetUserByEmail(ctx context.Context, email string) (user entities.User, err error)
	RepositoryDB
}

type RepositoryDB interface {
	Begin(ctx context.Context) (tx *sql.Tx, err error)
	Rollback(ctx context.Context, tx *sql.Tx) (err error)
	Commit(ctx context.Context, tx *sql.Tx) (err error)
}

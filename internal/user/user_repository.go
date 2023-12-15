package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User, profile *Profile) (*User, error) {
	var lastInsertedID string
	query := "INSERT INTO users (email, password, role) VALUES($1, $2, $3) returning id"

	err := r.db.QueryRowContext(ctx, query, user.Email, user.Password, user.Role).Scan(&lastInsertedID)
	if err != nil {
		return &User{}, err
	}

	query = "INSERT INTO profile (name, gender, users_id) VALUES($1, $2, $3) returning name"
	if _, err := r.db.ExecContext(ctx, query, profile.Name, profile.Gender, lastInsertedID); err != nil {
		return nil, err
	}

	parseUUID, _ := uuid.Parse(lastInsertedID)

	user.ID = parseUUID

	return user, nil
}

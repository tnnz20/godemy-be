package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type DBTX interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User, profile *Profile) (*User, *Profile, error) {
	var lastInsertedID uuid.UUID

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	defer tx.Rollback()

	query := "INSERT INTO users (email, password, role) VALUES($1, $2, $3) returning id"

	if err = tx.QueryRowContext(ctx, query, user.Email,
		user.Password, user.Role).Scan(&lastInsertedID); err != nil {
		return nil, nil, err
	}

	query = "INSERT INTO profile (name, gender, users_id) VALUES($1, $2, $3)"
	if _, err = tx.ExecContext(ctx, query, profile.Name,
		profile.Gender, lastInsertedID); err != nil {
		return nil, nil, err
	}

	// parseUUID, _ := uuid.Parse(lastInsertedID)
	user.ID = lastInsertedID

	if err = tx.Commit(); err != nil {
		return nil, nil, err
	}

	return user, profile, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email *string) (*User, error) {
	user := User{}
	query := "SELECT * FROM users WHERE email = $1"

	if err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email,
		&user.Password, &user.Role, &user.Created_at, &user.Updated_at); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserProfileByUserId(ctx context.Context, userId *uuid.UUID) (*User, *Profile, error) {
	user := User{}
	profile := Profile{}

	query := `SELECT u.id, u.email, u.role, u.created_at, u.updated_at, 
			p.name, p.gender, p.profile_img FROM users as u 
			JOIN profile as p ON u.id = p.users_id WHERE u.id = $1`

	if err := r.db.QueryRowContext(ctx, query, userId).Scan(&user.ID, &user.Email, &user.Role, &user.Created_at,
		&user.Updated_at, &profile.Name, &profile.Gender, &profile.ProfileImg); err != nil {
		return nil, nil, err
	}

	return &user, &profile, nil
}

func (r *repository) InsertRoleStudent(ctx context.Context, userId *uuid.UUID) error {

	query := "INSERT INTO student(users_id) VALUES ($1)"
	if _, err := r.db.ExecContext(ctx, query, userId); err != nil {
		return err
	}

	return nil
}

func (r *repository) InsertRoleTeacher(ctx context.Context, userId *uuid.UUID) error {

	query := "INSERT INTO teacher(users_id) VALUES ($1)"
	if _, err := r.db.ExecContext(ctx, query, userId); err != nil {
		return err
	}

	return nil
}

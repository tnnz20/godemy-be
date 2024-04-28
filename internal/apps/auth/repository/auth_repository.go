package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/auth"
	"github.com/tnnz20/godemy-be/internal/apps/auth/entities"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) auth.Repository {
	return repository{
		db: db,
	}
}

// Transaction DB

// Begin implements Repository.
func (r repository) Begin(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = r.db.BeginTx(ctx, &sql.TxOptions{})
	return
}

// Commit implements Repository.
func (repository) Commit(ctx context.Context, tx *sql.Tx) (err error) {
	return tx.Commit()
}

// Rollback implements Repository.
func (repository) Rollback(ctx context.Context, tx *sql.Tx) (err error) {
	return tx.Rollback()
}

func (r repository) CreateUserWithTX(ctx context.Context, tx *sql.Tx, user entities.User) (id uuid.UUID, err error) {
	query := `
	INSERT INTO users (email, password, role) 
	VALUES ($1, $2, $3) 
	RETURNING id
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return uuid.Nil, err
	}

	defer stmt.Close()

	err = tx.QueryRowContext(ctx, query, user.Email, user.Password, user.Role).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r repository) CreateProfileWithTX(ctx context.Context, tx *sql.Tx, profile entities.Profile) (err error) {
	query := `
	INSERT INTO profile (name, users_id) 
	VALUES ($1, $2)
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, profile.Name, profile.UserID)
	return
}

func (r repository) InsertUserRoleWithTX(ctx context.Context, tx *sql.Tx, user entities.User) (err error) {
	switch user.Role {
	case "student":
		query := `
		INSERT INTO student (users_id)
		VALUES ($1)
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}

		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, user.ID)
		return err

	case "teacher":
		query := `
		INSERT INTO teacher (users_id)
		VALUES ($1)
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}

		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, user.ID)
		return err
	}
	return
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (user entities.User, err error) {
	query := `
	SELECT id, email, password, role
	FROM users
	WHERE email = $1
	`

	err = r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		return
	}
	return
}

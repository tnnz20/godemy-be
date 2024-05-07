package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/users"
	"github.com/tnnz20/godemy-be/internal/apps/users/entities"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) users.Repository {
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

// CreateUsersWithTx implements user repository to create new users.
func (r repository) CreateUsersWithTX(ctx context.Context, tx *sql.Tx, user entities.Users) (id uuid.UUID, err error) {
	query := `
	INSERT INTO users (id, email, password, name, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return uuid.Nil, err
	}

	defer stmt.Close()

	err = tx.QueryRowContext(ctx, query,
		user.ID,
		user.Email,
		user.Password,
		user.Name,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r repository) InsertUsersRoleWithTX(ctx context.Context, tx *sql.Tx, role entities.Roles) (err error) {
	query := `
		INSERT INTO roles (users_id, role)
		VALUES ($1, $2)
		`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, role.UsersId, role.Role)
	return err

}

func (r repository) GetUserByEmail(ctx context.Context, email string) (user entities.Users, err error) {

	var (
		nullableDate       sql.NullTime
		nullableAddress    sql.NullString
		nullableGender     sql.NullString
		nullableProfileImg sql.NullString
	)

	query := `
	SELECT id, email, password, name, date, address, gender, profile_img, created_at, updated_at
	FROM users
	WHERE email = $1
	`

	err = r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&nullableDate,
		&nullableAddress,
		&nullableGender,
		&nullableProfileImg,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		return
	}

	if nullableDate.Valid {
		user.Date = nullableDate.Time
	}

	if nullableAddress.Valid {
		user.Address = nullableAddress.String
	}

	if nullableGender.Valid {
		user.Gender = nullableGender.String
	}

	if nullableProfileImg.Valid {
		user.ProfileImg = nullableProfileImg.String
	}

	return
}

func (r repository) GetRoleByUserID(ctx context.Context, userID uuid.UUID) (role entities.Roles, err error) {
	query := `
	SELECT users_id, role
	FROM roles
	WHERE users_id = $1
	`

	err = r.db.QueryRowContext(ctx, query, userID).Scan(
		&role.UsersId,
		&role.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		return
	}
	return
}

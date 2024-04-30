package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tnnz20/godemy-be/internal/apps/profile"
	"github.com/tnnz20/godemy-be/internal/apps/profile/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) profile.Repository {
	return &repository{
		db: db,
	}
}

func (r repository) FindProfileByUserId(ctx context.Context, userId uuid.UUID) (profile entities.Profile, err error) {
	query := `
	SELECT name, date, address, gender, profile_img, created_at, updated_at
	FROM profile
	WHERE users_id = $1
	`

	var (
		nullableDate       sql.NullTime
		nullableAddress    sql.NullString
		nullableGender     sql.NullString
		nullableProfileImg sql.NullString
	)
	err = r.db.QueryRowContext(ctx, query, userId).Scan(
		&profile.Name,
		&nullableDate,
		&nullableAddress,
		&nullableGender,
		&nullableProfileImg,
		&profile.CreatedAt,
		&profile.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errs.ErrUserNotFound
			return
		}
		return
	}

	if nullableDate.Valid {
		profile.Date = nullableDate.Time
	}

	if nullableAddress.Valid {
		profile.Address = nullableAddress.String
	}

	if nullableGender.Valid {
		profile.Gender = nullableGender.String
	}

	if nullableProfileImg.Valid {
		profile.ProfileImg = nullableProfileImg.String
	}

	return
}

package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tnnz20/godemy-be/internal/apps/auth"
	"github.com/tnnz20/godemy-be/internal/apps/auth/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type SvcRepository interface {
	auth.Repository
}

type service struct {
	repo SvcRepository
}

func newService(repo SvcRepository) auth.Service {
	return service{
		repo: repo,
	}
}

func (s service) Register(ctx context.Context, req entities.RegisterPayload) (err error) {

	regisUser := entities.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := regisUser.Validate(); err != nil {
		return err
	}

	// Check if email already exists
	user, err := s.repo.GetUserByEmail(ctx, regisUser.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return
		}
	}

	if user.IsEmailAlreadyExists() {
		return errs.ErrEmailAlreadyExists
	}

	// Begin transaction
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return err
	}

	// Rollback transaction if error
	defer s.repo.Rollback(ctx, tx)

	// Create user
	id, err := s.repo.CreateUserWithTX(ctx, tx, regisUser)
	if err != nil {
		return err
	}

	// Create profile
	regisProfile := entities.Profile{
		Name:   req.Name,
		UserID: id,
	}

	if err := regisProfile.Validate(); err != nil {
		return err
	}

	err = s.repo.CreateProfileWithTX(ctx, tx, regisProfile)
	if err != nil {
		return err
	}

	// Create role
	regisRole := entities.User{
		ID:   id,
		Role: req.Role,
	}
	err = s.repo.InsertUserRoleWithTX(ctx, tx, regisRole)
	if err != nil {
		return err
	}

	if err = s.repo.Commit(ctx, tx); err != nil {
		return
	}
	return
}

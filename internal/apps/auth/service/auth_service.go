package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tnnz20/godemy-be/internal/apps/auth"
	"github.com/tnnz20/godemy-be/internal/apps/auth/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type service struct {
	repo        auth.Repository
	secretToken string
}

func NewService(repo auth.Repository, secret string) auth.Service {
	return service{
		repo:        repo,
		secretToken: secret,
	}
}

func (s service) Register(ctx context.Context, req entities.RegisterPayload) (err error) {

	NewUser := entities.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := NewUser.Validate(); err != nil {
		return err
	}

	// Check if email already exists
	user, err := s.repo.GetUserByEmail(ctx, NewUser.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return
		}
	}

	if user.IsEmailAlreadyExists() {
		return errs.ErrEmailAlreadyExists
	}

	// Hash password
	if err := NewUser.HashingPassword(); err != nil {
		return err
	}

	// Begin transaction
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return err
	}

	// Rollback transaction if error
	defer s.repo.Rollback(ctx, tx)

	// Create user
	id, err := s.repo.CreateUserWithTX(ctx, tx, NewUser)
	if err != nil {
		return err
	}

	// Create profile
	NewProfile := entities.Profile{
		Name:   req.Name,
		UserID: id,
	}

	if err := NewProfile.Validate(); err != nil {
		return err
	}

	err = s.repo.CreateProfileWithTX(ctx, tx, NewProfile)
	if err != nil {
		return err
	}

	// Create role
	NewRole := entities.User{
		ID:   id,
		Role: NewUser.Role,
	}
	err = s.repo.InsertUserRoleWithTX(ctx, tx, NewRole)
	if err != nil {
		return err
	}

	if err = s.repo.Commit(ctx, tx); err != nil {
		return
	}
	return
}

func (s service) Login(ctx context.Context, req entities.LoginPayload) (res entities.LoginResponse, err error) {
	NewUserLogin := entities.User{
		Email:    req.Email,
		Password: req.Password,
	}

	// Validate email and password
	if err := NewUserLogin.ValidateEmail(); err != nil {
		return entities.LoginResponse{}, err
	}

	if err := NewUserLogin.ValidatePassword(); err != nil {
		return entities.LoginResponse{}, err
	}

	// Get user by email
	user, err := s.repo.GetUserByEmail(ctx, NewUserLogin.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.LoginResponse{}, errs.ErrEmailNotFound
		}
		return entities.LoginResponse{}, err
	}

	// Compare password
	if err := NewUserLogin.VerifyPasswordFromPlain(user.Password); err != nil {
		err = errs.ErrWrongPassword
		return entities.LoginResponse{}, err
	}

	// Generate token
	token, err := user.GenerateToken(s.secretToken)
	if err != nil {
		return entities.LoginResponse{}, err
	}

	res = entities.LoginResponse{
		Token: token,
	}

	return
}

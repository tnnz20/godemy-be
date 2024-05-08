package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tnnz20/godemy-be/internal/apps/users"
	"github.com/tnnz20/godemy-be/internal/apps/users/entities"
	"github.com/tnnz20/godemy-be/pkg/errs"
)

type service struct {
	users.Repository
	secretToken string
}

func NewService(userRepo users.Repository, secret string) users.Service {
	return service{
		Repository:  userRepo,
		secretToken: secret,
	}
}

func (s service) Register(ctx context.Context, req entities.RegisterPayload) (err error) {

	NewUser := entities.NewUsersRegister(req.Email, req.Password, req.Name)

	// Validate user payload
	if err := NewUser.ValidateAuth(); err != nil {
		return err
	}

	// Check if email already exists
	user, err := s.Repository.GetUserByEmail(ctx, NewUser.Email)
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
	tx, err := s.Repository.Begin(ctx)
	if err != nil {
		return err
	}

	// Rollback transaction if error
	defer s.Repository.Rollback(ctx, tx)

	// Create user
	id, err := s.Repository.CreateUsersWithTX(ctx, tx, NewUser)
	if err != nil {
		return err
	}

	// Create role
	NewRole := entities.NewRoles(id, req.Role)

	if err := NewRole.ValidateRole(); err != nil {
		return err
	}

	err = s.Repository.InsertUsersRoleWithTX(ctx, tx, NewRole)
	if err != nil {
		return err
	}

	if err = s.Repository.Commit(ctx, tx); err != nil {
		return
	}
	return
}

func (s service) Login(ctx context.Context, req entities.LoginPayload) (res entities.LoginResponse, err error) {
	NewUserLogin := entities.NewUsersLogin(req.Email, req.Password)

	// Validate email and password
	if err := NewUserLogin.ValidateEmail(); err != nil {
		return entities.LoginResponse{}, err
	}

	if err := NewUserLogin.ValidatePassword(); err != nil {
		return entities.LoginResponse{}, err
	}

	// Get user by email
	user, err := s.Repository.GetUserByEmail(ctx, NewUserLogin.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.LoginResponse{}, errs.ErrEmailNotFound
		}
		return entities.LoginResponse{}, err
	}

	role, err := s.Repository.GetRoleByUserID(ctx, user.ID)
	if err != nil {
		return entities.LoginResponse{}, err
	}

	// Compare password
	if err := NewUserLogin.VerifyPasswordFromPlain(user.Password); err != nil {
		err = errs.ErrWrongPassword
		return entities.LoginResponse{}, err
	}

	// Generate token
	token, err := user.GenerateToken(role.Role, s.secretToken)
	if err != nil {
		return entities.LoginResponse{}, err
	}

	res = entities.LoginResponse{
		Token: token,
	}

	return
}

func (s service) GetUser(ctx context.Context, req entities.GetUserPayload) (res entities.UserResponse, err error) {
	user, err := s.Repository.GetUserByUserId(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errs.ErrUserNotFound
			return
		}
		return
	}

	res = entities.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Date:       user.Date,
		Address:    user.Address,
		Gender:     user.Gender,
		ProfileImg: user.ProfileImg,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	return
}

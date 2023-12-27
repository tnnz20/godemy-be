package user

import (
	"context"
	"time"

	"github.com/tnnz20/godemy-be/util"
)

type service struct {
	UserRepository Repository
	timeout        time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		UserRepository: repository,
		timeout:        time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	profile := &Profile{
		Name:   req.Name,
		Gender: req.Gender,
	}

	userRes, profileRes, err := s.UserRepository.CreateUser(ctx, user, profile)
	if err != nil {
		return nil, err
	}

	switch req.Role {
	case "teacher":
		if err := s.UserRepository.InsertRoleTeacher(ctx, &userRes.ID); err != nil {
			return nil, err
		}
	case "student":
		if err := s.UserRepository.InsertRoleStudent(ctx, &userRes.ID); err != nil {
			return nil, err
		}
	}

	res := &CreateUserResponse{
		ID:    userRes.ID,
		Email: userRes.Email,
		Name:  profileRes.Name,
		Role:  user.Role,
	}

	return res, nil
}

func (s *service) GetUserProfileById(c context.Context, req *GetUserProfileByIdRequest) (*GetUserProfileByIdResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	userRes, profileRes, err := s.UserRepository.GetUserProfileById(ctx, &req.ID)
	if err != nil {
		return nil, err
	}

	res := &GetUserProfileByIdResponse{
		ID:     userRes.ID,
		Email:  userRes.Email,
		Role:   userRes.Role,
		Name:   profileRes.Name,
		Gender: profileRes.Gender,
	}

	return res, nil
}

func (s *service) GetUserByEmail(c context.Context, email *string) (*User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	res := &User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Email,
		Role:     user.Role,
	}

	return res, nil
}

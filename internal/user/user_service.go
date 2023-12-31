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

	userReq := &User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	profileReq := &Profile{
		Name:   req.Name,
		Gender: req.Gender,
	}

	userRes, profileRes, err := s.UserRepository.CreateUser(ctx, userReq, profileReq)
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
		Role:  userRes.Role,
	}

	return res, nil
}

func (s *service) GetUserProfileByUserId(c context.Context, req *GetUserProfileByUserIdRequest) (*GetUserProfileByUserIdResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	userRes, profileRes, err := s.UserRepository.GetUserProfileByUserId(ctx, &req.UserId)
	if err != nil {
		return nil, err
	}

	res := &GetUserProfileByUserIdResponse{
		ID:         userRes.ID,
		Email:      userRes.Email,
		Role:       userRes.Role,
		Name:       profileRes.Name,
		Gender:     profileRes.Gender,
		ProfileImg: profileRes.ProfileImg,
	}

	return res, nil
}

func (s *service) GetUserByEmail(c context.Context, req *GetUserByEmailRequest) (*GetUserByEmailResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.UserRepository.GetUserByEmail(ctx, &req.Email)
	if err != nil {
		return nil, err
	}

	res := &GetUserByEmailResponse{
		ID:         user.ID,
		Email:      user.Email,
		Password:   user.Password,
		Role:       user.Role,
		Created_at: user.Created_at,
		Updated_at: user.Deleted_at,
	}

	return res, nil
}

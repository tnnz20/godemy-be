package user

import (
	"context"
	"time"

	"github.com/tnnz20/godemy-be/util"
)

// const (
// 	secretKey = "secret"
// )

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq, isTeacher bool) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	var userRole string
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	if isTeacher {
		userRole = "teacher"
	} else {
		userRole = "student"
	}

	u := &User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     userRole,
	}

	p := &Profile{
		Name:   req.Name,
		Gender: req.Gender,
	}

	r, err := s.Repository.CreateUser(ctx, u, p)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:    r.ID,
		Email: r.Email,
	}

	return res, nil
}

package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tnnz20/godemy-be/util"
)

type service struct {
	Repository
	timeout time.Duration
	jwtKey  *string
}

func NewService(repository Repository, jwtKey *string) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
		jwtKey,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq, isTeacher bool) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	var userRoleReq string

	if res, err := s.Repository.GetUserByEmail(ctx, &req.Email); err == nil && res.Email == req.Email {
		return nil, fmt.Errorf("Email %s already exists", req.Email)
	}

	if isTeacher {
		userRoleReq = "teacher"
	} else {
		userRoleReq = "student"
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     userRoleReq,
	}

	profile := &Profile{
		Name:   req.Name,
		Gender: req.Gender,
	}

	userRes, profileRes, err := s.Repository.CreateUser(ctx, user, profile)
	if err != nil {
		return nil, err
	}

	switch userRoleReq {
	case "teacher":
		if err := s.Repository.InsertRoleTeacher(ctx, &userRes.ID); err != nil {
			return nil, err
		}
	case "student":
		if err := s.Repository.InsertRoleStudent(ctx, &userRes.ID); err != nil {
			return nil, err
		}
	}

	res := &CreateUserRes{
		ID:    userRes.ID,
		Email: userRes.Email,
		Name:  profileRes.Name,
		Role:  userRoleReq,
	}

	return res, nil
}

func (s *service) GetUserProfileById(c context.Context, req *GetUserProfileByIdReq) (*GetUserProfileByIdRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	userRes, profileRes, err := s.Repository.GetUserProfileById(ctx, &req.ID)
	if err != nil {
		return &GetUserProfileByIdRes{}, err
	}

	res := &GetUserProfileByIdRes{
		ID:     userRes.ID,
		Email:  userRes.Email,
		Role:   userRes.Role,
		Name:   profileRes.Name,
		Gender: profileRes.Gender,
	}

	return res, nil
}

func (s *service) SignIn(c context.Context, req *SignInReq) (*SignInRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.Repository.GetUserByEmail(ctx, &req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with email %s not found", req.Email)
		}
	}

	if err := util.CheckPassword(req.Password, user.Password); err != nil {
		return nil, fmt.Errorf("Invalid Password")
	}

	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := t.SignedString([]byte(*s.jwtKey))
	if err != nil {
		return nil, err
	}

	res := &SignInRes{
		Token: token,
	}
	return res, nil
}

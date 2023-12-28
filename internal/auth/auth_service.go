package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/tnnz20/godemy-be/internal/user"
	"github.com/tnnz20/godemy-be/util"
)

type service struct {
	UserRepository user.Repository
	timeout        time.Duration
	jwtKey         *string
}

func NewService(repository user.Repository, key *string) Service {
	return &service{
		UserRepository: repository,
		timeout:        time.Duration(2) * time.Second,
		jwtKey:         key,
	}
}

func (s *service) SignIn(c context.Context, req *AuthRequest) (*AuthResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.UserRepository.GetUserByEmail(ctx, &req.Email)
	if err != nil {
		return nil, err
	}

	if err := util.CheckPassword(req.Password, user.Password); err != nil {
		return nil, fmt.Errorf("Invalid Password.")
	}

	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := t.SignedString([]byte(*s.jwtKey))
	if err != nil {
		return nil, err
	}

	res := &AuthResponse{
		Token: token,
	}
	return res, nil
}

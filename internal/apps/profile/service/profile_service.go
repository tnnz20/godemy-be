package service

import (
	"context"

	"github.com/tnnz20/godemy-be/internal/apps/profile"
	"github.com/tnnz20/godemy-be/internal/apps/profile/entities"
)

type service struct {
	profile.Repository
	secret string
}

// TODO: change time to unix timestamp
func NewService(profileRepo profile.Repository, secret string) profile.Service {
	return &service{
		Repository: profileRepo,
		secret:     secret,
	}
}

func (s *service) GetProfileByUserId(ctx context.Context, req entities.GetProfileByUserIdRequest) (res *entities.GetProfileByUserIdResponse, err error) {
	user := entities.Profile{
		UserID: req.UserId,
	}

	if err := user.ValidateUserID(); err != nil {
		return nil, err
	}

	profileUser, err := s.Repository.FindProfileByUserId(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	res = &entities.GetProfileByUserIdResponse{
		Name:       profileUser.Name,
		Date:       profileUser.Date,
		Address:    profileUser.Address,
		Gender:     profileUser.Gender,
		ProfileImg: profileUser.ProfileImg,
		CreatedAt:  profileUser.CreatedAt,
		UpdatedAt:  profileUser.UpdatedAt,
	}

	return res, nil

}

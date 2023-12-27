package teacher

import (
	"context"
	"time"
)

type service struct {
	TeacherRepository Repository
	timeout           time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		TeacherRepository: repository,
		timeout:           time.Duration(2) * time.Second,
	}
}

func (s *service) CreateClass(c context.Context, req *CreateClassRequest) (*CreateClassResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classRequest := &Class{
		TeacherId: req.TeacherId,
		ClassName: req.ClassName,
	}
	classResponse, err := s.TeacherRepository.CreateClass(ctx, classRequest)
	if err != nil {
		return nil, err
	}

	res := &CreateClassResponse{
		ID:        classResponse.ID,
		TeacherId: classResponse.TeacherId,
		ClassName: classResponse.ClassName,
	}

	return res, nil
}

func (s *service) GetTeacherIdByUserId(c context.Context, req *GetTeacherIdByUserIdRequest) (*GetTeacherIdByUserIdResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	teacher, err := s.TeacherRepository.GetTeacherIdByUserId(ctx, &req.ID)
	if err != nil {
		return nil, err
	}

	res := &GetTeacherIdByUserIdResponse{
		ID:     teacher.ID,
		UserId: teacher.UserId,
	}

	return res, nil
}

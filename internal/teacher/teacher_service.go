package teacher

import (
	"context"
	"fmt"
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

func (s *service) GetTeacherIdByUserId(c context.Context, req *GetTeacherIdByUserIdRequest) (*GetTeacherIdByUserIdResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	teacher, err := s.TeacherRepository.GetTeacherIdByUserId(ctx, &req.UserId)
	if err != nil {
		return nil, err
	}

	res := &GetTeacherIdByUserIdResponse{
		ID:     teacher.ID,
		UserId: teacher.UserId,
	}

	return res, nil
}

func (s *service) GetAllClassByTeacherId(c context.Context, req *GetClassByTeacherIdRequest) (*[]GetClassByTeacherIdResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.TeacherRepository.GetAllClassByTeacherId(ctx, &req.TeacherId)
	if err != nil {
		return nil, err
	} else if len(*classes) == 0 {
		return nil, fmt.Errorf("null")
	}

	var response []GetClassByTeacherIdResponse
	for _, class := range *classes {
		res := GetClassByTeacherIdResponse{
			ID:        class.ID,
			TeacherId: class.TeacherId,
			ClassName: class.ClassName,
		}
		response = append(response, res)
	}

	return &response, nil
}

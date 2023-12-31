package student

import (
	"context"
	"time"
)

type service struct {
	StudentRepository Repository
	timeout           time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		StudentRepository: repository,
		timeout:           time.Duration(2) * time.Second,
	}
}

func (s *service) GetStudentByUserId(c context.Context, req *GetStudentByUserIdRequest) (*GetStudentByUserIdResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	studentRes, err := s.StudentRepository.GetStudentByUserId(ctx, &req.UsersId)
	if err != nil {
		return nil, err
	}

	res := &GetStudentByUserIdResponse{
		ID:        studentRes.ID,
		UsersId:   studentRes.UsersId,
		ClassId:   studentRes.ClassId,
		Threshold: studentRes.Threshold,
	}

	return res, nil
}

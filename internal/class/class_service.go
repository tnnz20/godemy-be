package class

import (
	"context"
	"fmt"
	"time"
)

type service struct {
	ClassRepository Repository
	timeout         time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		ClassRepository: repository,
		timeout:         time.Duration(2) * time.Second,
	}
}

func (s *service) CreateClass(c context.Context, req *CreateClassRequest) (*CreateClassResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classRequest := &Class{
		TeacherId: req.TeacherId,
		ClassName: req.ClassName,
	}
	classResponse, err := s.ClassRepository.CreateClass(ctx, classRequest)
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

func (s *service) GetAllClass(c context.Context) (*[]GetAllClassResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.ClassRepository.GetAllClass(ctx)
	if err != nil {
		return nil, err
	} else if len(*classes) == 0 {
		return nil, fmt.Errorf("null")
	}

	var response []GetAllClassResponse
	for _, class := range *classes {
		res := GetAllClassResponse{
			ID:        class.ID,
			TeacherId: class.TeacherId,
			ClassName: class.ClassName,
		}
		response = append(response, res)
	}

	return &response, nil
}

func (s *service) UpdateStudentClass(c context.Context, req *UpdateStudentClassRequest) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	if err := s.ClassRepository.UpdateStudentClass(ctx, &req.ClassId, &req.StudentId); err != nil {
		return err
	}

	return nil
}

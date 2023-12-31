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

func (s *service) IncrementThreshold(c context.Context, req *IncrementThresholdRequest) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	if err := s.StudentRepository.IncrementThreshold(ctx, &req.UsersId, &req.Threshold); err != nil {
		return err
	}

	return nil
}

func (s *service) InsertAssessment(c context.Context, req *InsertAssessmentRequest) (*InsertAssessmentResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	assessmentReq := &Assessment{
		AssessmentValue: req.AssessmentValue,
		CodeAssessment:  req.CodeAssessment,
	}
	user, assessment, err := s.StudentRepository.InsertAssessment(ctx, &req.UsersId, assessmentReq)
	if err != nil {
		return nil, err
	}

	res := &InsertAssessmentResponse{
		StudentId:       user.ID,
		AssessmentValue: assessment.AssessmentValue,
		CodeAssessment:  assessment.CodeAssessment,
	}

	return res, nil
}

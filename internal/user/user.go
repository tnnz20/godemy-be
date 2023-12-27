package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Ordered_at time.Time `json:"ordered_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Profile struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	UserId     uuid.UUID `json:"user_id"`
	ProfileImg string    `json:"profile_img"`
}

type Student struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	ClassId   uuid.UUID `json:"class_id"`
	Threshold int       `json:"threshold"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required,oneof='Laki-Laki' 'Perempuan'" `
	Role     string `json:"role" validate:"required,oneof='student' 'teacher'"`
}

type CreateUserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}

type GetUserProfileByIdRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid4"`
}
type GetUserProfileByIdResponse struct {
	ID     uuid.UUID `json:"id"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	Name   string    `json:"name"`
	Gender string    `json:"gender"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User, profile *Profile) (*User, *Profile, error)
	GetUserByEmail(ctx context.Context, email *string) (*User, error)
	GetUserProfileById(ctx context.Context, id *uuid.UUID) (*User, *Profile, error)
	InsertRoleStudent(ctx context.Context, userId *uuid.UUID) error
	InsertRoleTeacher(ctx context.Context, userId *uuid.UUID) error
}

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
	GetUserProfileById(ctx context.Context, req *GetUserProfileByIdRequest) (*GetUserProfileByIdResponse, error)
	GetUserByEmail(ctx context.Context, email *string) (*User, error)
}

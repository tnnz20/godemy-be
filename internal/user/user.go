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

type CreateUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required,oneof='Laki-Laki' 'Perempuan'" `
}

type CreateUserRes struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User, profile *Profile) (*User, error)
	// GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq, isTeacher bool) (*CreateUserRes, error)
}

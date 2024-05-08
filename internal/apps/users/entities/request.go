package entities

import "github.com/google/uuid"

type RegisterPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserPayload struct {
	ID uuid.UUID `json:"id"`
}

type UpdateUserPayload struct {
	ID uuid.UUID `json:"id"`
}

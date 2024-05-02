package entities

import (
	"github.com/google/uuid"
)

type Teacher struct {
	ID     uuid.UUID
	UserId uuid.UUID
}

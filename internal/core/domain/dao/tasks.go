package entities

import (
	"time"

	"github.com/google/uuid"
)

type Tasks struct {
	TaskID      uuid.UUID
	WorkspaceID uuid.UUID
	UserID      uuid.UUID
	Title       string
	Status      string
	Priority    string
	CreatedAt   time.Time
}

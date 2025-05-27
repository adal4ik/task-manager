package entities

import (
	"time"

	"github.com/google/uuid"
)

type Workspace struct {
	WorkspaceID uuid.UUID
	Name        string
	IsPrivate   bool
	CreatedAt   time.Time
}

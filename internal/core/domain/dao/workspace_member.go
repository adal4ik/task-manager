package dao

import (
	"time"

	"github.com/google/uuid"
)

type WorkspaceMember struct {
	UserID      uuid.UUID
	WorkspaceID uuid.UUID
	Role        string
	JoinedAt    time.Time
}

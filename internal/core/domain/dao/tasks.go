package dao

import (
	"time"

	"github.com/google/uuid"
)

type Tasks struct {
	TaskID    uuid.UUID
	UserID    uuid.UUID
	Title     string
	Status    string
	Priority  string
	DueDate   time.Time
	CreatedAt time.Time
}

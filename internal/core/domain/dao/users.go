package dao

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	UserID    uuid.UUID
	Login     string
	Email     string
	Password  string
	CreatedAt time.Time
}

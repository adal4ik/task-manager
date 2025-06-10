package dao

import (
	"time"
)

type Tasks struct {
	TaskID    string
	UserID    string
	Title     string
	Status    string
	Priority  string
	DueDate   time.Time
	CreatedAt time.Time
}

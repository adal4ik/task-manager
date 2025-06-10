package dto

import (
	"fmt"
	"task-manager/internal/core/domain/dao"
	"time"
)

type Task struct {
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	Priority  string `json:"priority"`
	DueDate   string `json:"due_date"`   // ISO 8601 format
	CreatedAt string `json:"created_at"` // ISO 8601 format
}

func TaskToDao(task Task) (dao.Tasks, error) {
	var createdAt time.Time
	var err error
	if task.CreatedAt == "" {
		createdAt = time.Now().UTC()
	} else {
		createdAt, err = time.Parse(time.RFC3339, task.CreatedAt)
		if err != nil {
			return dao.Tasks{}, fmt.Errorf("invalid created_at: %w", err)
		}
	}

	dueDate, err := time.Parse(time.RFC3339, task.DueDate)
	if err != nil {
		return dao.Tasks{}, fmt.Errorf("invalid due_date: %w", err)
	}

	return dao.Tasks{
		UserID:    task.UserID,
		Title:     task.Title,
		Status:    task.Status,
		Priority:  task.Priority,
		DueDate:   dueDate,
		CreatedAt: createdAt,
	}, nil
}

func DaoToTask(task dao.Tasks) Task {
	return Task{
		UserID:    task.UserID,
		Title:     task.Title,
		Status:    task.Status,
		Priority:  task.Priority,
		DueDate:   task.DueDate.Format(time.RFC3339),
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
	}
}

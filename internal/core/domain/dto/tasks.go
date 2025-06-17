package dto

import (
	"fmt"
	"task-manager/internal/core/domain/dao"
	"time"
)

type UpdateStatus struct {
	Status string `json:"status"`
}

type Task struct {
	TaskID    string     `json:"task_id,omitempty"`
	UserID    string     `json:"user_id,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Status    *string    `json:"status,omitempty"`
	Priority  *string    `json:"priority,omitempty"`
	DueDate   *time.Time `json:"due_date,omitempty"`
	CreatedAt string     `json:"created_at,omitempty"`
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

	return dao.Tasks{
		TaskID:    task.TaskID,
		UserID:    task.UserID,
		Title:     task.Title,
		Status:    task.Status,
		Priority:  task.Priority,
		DueDate:   task.DueDate,
		CreatedAt: createdAt,
	}, nil
}

func DaoToTask(task dao.Tasks) Task {
	var dueDate *time.Time
	if task.DueDate != nil {
		dueDate = task.DueDate
	}

	return Task{
		TaskID:    task.TaskID,
		UserID:    task.UserID,
		Title:     task.Title,
		Status:    task.Status,
		Priority:  task.Priority,
		DueDate:   dueDate,
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
	}
}

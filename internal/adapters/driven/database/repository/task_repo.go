package repository

import (
	"context"
	"database/sql"
	"task-manager/internal/core/domain/dto"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) CreateTask(ctx context.Context, task dto.Task) error {
	query := `INSERT INTO tasks (user_id, title, status, priority, due_date) VALUES ($1, $2, $3, $4, $5)`
	_, err := t.db.ExecContext(ctx, query, task.UserID, task.Title, task.Status, task.Priority, task.DueDate)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"
	"database/sql"
	"task-manager/internal/core/domain/dao"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) CreateTask(ctx context.Context, task dao.Tasks) error {
	query := `INSERT INTO tasks (user_id, title, status, priority, due_date) VALUES ($1, $2, $3, $4, $5)`
	_, err := t.db.ExecContext(ctx, query, task.UserID, task.Title, task.Status, task.Priority, task.DueDate)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) GetTasks(ctx context.Context, userID string) ([]dao.Tasks, error) {
	query := `SELECT user_id, title, status, priority, due_date, created_at FROM tasks WHERE user_id = $1`
	rows, err := t.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dao.Tasks
	for rows.Next() {
		var task dao.Tasks
		if err := rows.Scan(&task.UserID, &task.Title, &task.Status, &task.Priority, &task.DueDate, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

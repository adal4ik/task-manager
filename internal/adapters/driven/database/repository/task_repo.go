package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"task-manager/internal/core/domain/dao"
	"task-manager/internal/utils"
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
	query := `SELECT task_id,user_id, title, status, priority, due_date, created_at FROM tasks WHERE user_id = $1`
	rows, err := t.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dao.Tasks
	for rows.Next() {
		var task dao.Tasks
		if err := rows.Scan(&task.TaskID, &task.UserID, &task.Title, &task.Status, &task.Priority, &task.DueDate, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskRepository) UpdateTask(ctx context.Context, task dao.Tasks, taskID string) error {
	setParts := []string{}
	args := []interface{}{}
	argIdx := 1

	if task.Title != nil && *task.Title != "" {
		setParts = append(setParts, fmt.Sprintf("title = $%d", argIdx))
		args = append(args, task.Title)
		argIdx++
	}
	if task.Priority != nil && *task.Priority != "" {
		setParts = append(setParts, fmt.Sprintf("priority = $%d", argIdx))
		args = append(args, task.Priority)
		argIdx++
	}
	if task.DueDate != nil && !task.DueDate.IsZero() {
		setParts = append(setParts, fmt.Sprintf("due_date = $%d", argIdx))
		args = append(args, task.DueDate)
		argIdx++
	}

	if len(setParts) == 0 {
		return nil
	}

	args = append(args, task.UserID, taskID)
	query := fmt.Sprintf(`UPDATE tasks SET %s WHERE user_id = $%d AND task_id = $%d`,
		strings.Join(setParts, ", "), argIdx, argIdx+1)

	_, err := t.db.ExecContext(ctx, query, args...)
	return err
}

func (t *TaskRepository) DeleteTask(ctx context.Context, userID string, taskID string) error {
	query := `DELETE FROM tasks WHERE user_id = $1 AND task_id = $2`
	_, err := t.db.ExecContext(ctx, query, userID, taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.ErrNoRows
		}
		return err
	}
	return nil
}

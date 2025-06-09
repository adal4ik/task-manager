package dto

type Task struct {
	TaskID    string `json:"task_id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	Priority  string `json:"priority"`
	DueDate   string `json:"due_date"`   // ISO 8601 format
	CreatedAt string `json:"created_at"` // ISO 8601 format
}

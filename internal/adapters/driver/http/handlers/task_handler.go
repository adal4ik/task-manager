package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/internal/adapters/driver/http/middleware"
	"task-manager/internal/core/domain/dto"
	"task-manager/internal/core/interfaces/driven"
	"task-manager/internal/utils"
)

type TaskHandler struct {
	service driven.TasksDrivenInteface
	BaseHandler
}

func NewTaskHandler(service driven.TasksDrivenInteface, basehandler BaseHandler) *TaskHandler {
	return &TaskHandler{
		BaseHandler: basehandler,
		service:     service,
	}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var task dto.Task
	err := decoder.Decode(&task)
	if err != nil {
		t.handleError(w, req, http.StatusInternalServerError, "Failed to decode request body", err)
		return
	}
	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		t.handleError(w, req, http.StatusUnauthorized, "User ID not found in context", nil)
		return
	}
	task.UserID = userID
	err = t.service.CreateTask(req.Context(), task)
	if err != nil {
		t.handleError(w, req, http.StatusInternalServerError, "Failed to create task", err)
		return
	}
	t.logger.Info("Task created successfully", "task", task)
	resp := utils.APIResponse{
		Code:    http.StatusCreated,
		Message: "Task created successfully",
	}
	resp.Send(w)
}

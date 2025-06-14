package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"task-manager/internal/adapters/driver/http/middleware"
	"task-manager/internal/core/domain/dto"
	"task-manager/internal/core/interfaces/driver"
	"task-manager/internal/utils"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	service driver.TasksDriverInterface
	BaseHandler
}

func NewTaskHandler(service driver.TasksDriverInterface, basehandler BaseHandler) *TaskHandler {
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

func (t *TaskHandler) GetTasks(w http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		t.handleError(w, req, http.StatusUnauthorized, "User ID not found in context", nil)
		return
	}
	tasks, err := t.service.GetTasks(req.Context(), userID)
	if err != nil {
		t.handleError(w, req, http.StatusInternalServerError, "Failed to get tasks", err)
		return
	}
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		t.handleError(w, req, http.StatusInternalServerError, "Failed to marshal tasks", err)
		return
	}
	// json.NewEncoder(w).Encode(tasks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (t *TaskHandler) PatchTask(w http.ResponseWriter, req *http.Request) {
	taskID := chi.URLParam(req, "task_id")
	if taskID == "" {
		t.handleError(w, req, http.StatusBadRequest, "Task ID is required", nil)
		return
	}
	data := json.NewDecoder(req.Body)
	var task dto.Task
	err := data.Decode(&task)
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
	err = t.service.UpdateTask(req.Context(), task, taskID)
	if err != nil {
		t.handleError(w, req, http.StatusInternalServerError, "Failed to update task", err)
		return
	}
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, req *http.Request) {
	taskID := chi.URLParam(req, "task_id")
	if taskID == "" {
		t.handleError(w, req, http.StatusBadRequest, "Task ID is required", nil)
		return
	}
	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		t.handleError(w, req, http.StatusUnauthorized, "User ID not found in context", nil)
		return
	}
	err := t.service.DeleteTask(req.Context(), userID, taskID)
	if err != nil {
		if errors.Is(err, utils.ErrNoRows) {
			t.handleError(w, req, http.StatusBadRequest, "Task not found", err)
			return
		}
		t.handleError(w, req, http.StatusInternalServerError, "Something went wrong", err)
		return
	}
	resp := utils.APIResponse{
		Code:    204,
		Message: "Successfuly deleted",
	}
	resp.Send(w)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/internal/core/domain/dto"
	"task-manager/internal/core/interfaces/driver"
	"task-manager/internal/utils"
)

type LoginHandler struct {
	service driver.LoginDriverInterface
	BaseHandler
}

func NewLoginHandler(service driver.LoginDriverInterface, baseHandler BaseHandler) *LoginHandler {
	return &LoginHandler{
		service:     service,
		BaseHandler: baseHandler,
	}
}

func (l *LoginHandler) LoginUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login := req.FormValue("login")
	password := req.FormValue("password")
	if login == "" || password == "" {
		l.handleError(w, req, http.StatusBadRequest, "Missing required fields", nil)
		return
	}
	token, err := l.service.LoginUser(req.Context(), login, password)
	if err != nil {
		switch err {
		case utils.ErrNoRows:
			l.handleError(w, req, http.StatusNotFound, "User not found", err)
		case utils.ErrInvalidCredentials:
			l.handleError(w, req, http.StatusUnauthorized, "Invalid credentials", err)
		case utils.ErrMissingSecret:
			l.handleError(w, req, http.StatusInternalServerError, "Missing JWT secret", err)
		default:
			l.handleError(w, req, http.StatusInternalServerError, "Internal server error", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.LoginResponse{Token: token})
}

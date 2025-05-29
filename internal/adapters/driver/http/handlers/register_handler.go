package handlers

import (
	"net/http"
	"task-manager/internal/core/interfaces/driver"
	"task-manager/internal/utils"
)

type RegisterHandler struct {
	service driver.RegisterDriverInterface
	BaseHandler
}

func NewRegisterHandler(service driver.RegisterDriverInterface, basehandler BaseHandler) *RegisterHandler {
	return &RegisterHandler{
		service:     service,
		BaseHandler: basehandler,
	}
}

func (r *RegisterHandler) RegisterUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login := req.FormValue("login")
	password := req.FormValue("password")
	email := req.FormValue("email")
	if login == "" || password == "" || email == "" {
		r.handleError(w, req, http.StatusBadRequest, "Missing required fields", nil)
		return
	}
	err := r.service.RegisterUser(req.Context(), login, password, email)
	if err != nil {
		r.handleError(w, req, http.StatusInternalServerError, "Failed to register user", err)
		return
	}
	json := utils.APIResponse{
		Code:    http.StatusCreated,
		Message: "User registered successfully",
	}
	json.Send(w)
	return
}

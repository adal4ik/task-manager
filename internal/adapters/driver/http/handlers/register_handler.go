package handlers

import (
	"errors"
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

func (r *RegisterHandler) validateInput(w http.ResponseWriter, req *http.Request, login, password, email string) error {
	if login == "" || password == "" || email == "" {
		r.handleError(w, req, http.StatusBadRequest, "Missing required fields", nil)
		return errors.New("missing fields")
	}
	return nil
}

func (r *RegisterHandler) checkConflicts(w http.ResponseWriter, req *http.Request, login, email string) error {
	ctx := req.Context()

	exists, err := r.service.CheckEmailExists(ctx, email)
	if err != nil {
		r.handleError(w, req, http.StatusInternalServerError, "Failed to check email existence", err)
		return err
	}
	if exists {
		r.handleError(w, req, http.StatusConflict, "Email already exists", nil)
		return errors.New("email exists")
	}

	exists, err = r.service.CheckLoginExists(ctx, login)
	if err != nil {
		r.handleError(w, req, http.StatusInternalServerError, "Failed to check login existence", err)
		return err
	}
	if exists {
		r.handleError(w, req, http.StatusConflict, "Login already exists", nil)
		return errors.New("login exists")
	}

	return nil
}

func (r *RegisterHandler) RegisterUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login := req.FormValue("login")
	password := req.FormValue("password")
	email := req.FormValue("email")

	if err := r.validateInput(w, req, login, password, email); err != nil {
		return
	}

	if err := r.checkConflicts(w, req, login, email); err != nil {
		return
	}

	if err := r.service.RegisterUser(req.Context(), login, password, email); err != nil {
		r.handleError(w, req, http.StatusInternalServerError, "Failed to register user", err)
		return
	}
	resp := utils.APIResponse{
		Code:    http.StatusCreated,
		Message: "User registered successfully",
	}
	resp.Send(w)
}

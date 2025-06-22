package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"task-manager/internal/core/domain/dto"
	"task-manager/internal/core/interfaces/driver"
	"task-manager/internal/utils"
)

type AuthHandler struct {
	service driver.AuthDriverInterface
	BaseHandler
}

func NewAuthHandler(service driver.AuthDriverInterface, baseHandler BaseHandler) *AuthHandler {
	return &AuthHandler{
		service:     service,
		BaseHandler: baseHandler,
	}
}

func (a *AuthHandler) validateInput(w http.ResponseWriter, req *http.Request, login, password, email string) error {
	if login == "" || password == "" || email == "" {
		a.handleError(w, req, http.StatusBadRequest, "Missing required fields", nil)
		return errors.New("missing fields")
	}
	return nil
}

func (a *AuthHandler) LoginUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login := req.FormValue("login")
	password := req.FormValue("password")
	if login == "" || password == "" {
		a.handleError(w, req, http.StatusBadRequest, "Missing required fields", nil)
		return
	}
	token, err := a.service.LoginUser(req.Context(), login, password)
	if err != nil {
		switch err {
		case utils.ErrNoRows:
			a.handleError(w, req, http.StatusNotFound, "User not found", err)
		case utils.ErrInvalidCredentials:
			a.handleError(w, req, http.StatusUnauthorized, "Invalid credentials", err)
		case utils.ErrMissingSecret:
			a.handleError(w, req, http.StatusInternalServerError, "Missing JWT secret", err)
		default:
			a.handleError(w, req, http.StatusInternalServerError, "Internal server error", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.LoginResponse{Token: token})
}

func (a *AuthHandler) checkConflicts(w http.ResponseWriter, req *http.Request, login, email string) error {
	ctx := req.Context()

	exists, err := a.service.CheckEmailExists(ctx, email)
	if err != nil {
		a.handleError(w, req, http.StatusInternalServerError, "Failed to check email existence", err)
		return err
	}
	if exists {
		a.handleError(w, req, http.StatusConflict, "Email already exists", nil)
		return errors.New("email exists")
	}

	exists, err = a.service.CheckLoginExists(ctx, login)
	if err != nil {
		a.handleError(w, req, http.StatusInternalServerError, "Failed to check login existence", err)
		return err
	}
	if exists {
		a.handleError(w, req, http.StatusConflict, "Login already exists", nil)
		return errors.New("login exists")
	}

	return nil
}

func (a *AuthHandler) RegisterUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login := req.FormValue("login")
	password := req.FormValue("password")
	email := req.FormValue("email")

	if err := a.validateInput(w, req, login, password, email); err != nil {
		return
	}

	if err := a.checkConflicts(w, req, login, email); err != nil {
		return
	}

	if err := a.service.RegisterUser(req.Context(), login, password, email); err != nil {
		a.handleError(w, req, http.StatusInternalServerError, "Failed to register user", err)
		return
	}
	resp := utils.APIResponse{
		Code:    http.StatusCreated,
		Message: "User registered successfully",
	}
	resp.Send(w)
}

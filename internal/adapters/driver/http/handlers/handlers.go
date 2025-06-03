package handlers

import (
	"log/slog"
	"net/http"
	"task-manager/internal/core/service"
	"task-manager/internal/utils"
)

type BaseHandler struct {
	logger slog.Logger
}

func (b *BaseHandler) handleError(w http.ResponseWriter, r *http.Request, code int, message string, err error) {
	if err != nil {
		b.logger.Error(message, "error", err, "code", code, "url", r.URL.Path)
	} else {
		b.logger.Error(message, "code", code, "url", r.URL.Path)
	}

	jsonErr := utils.APIError{
		Code:     code,
		Message:  message,
		Resource: r.URL.Path,
	}
	jsonErr.Send(w)
}

func NewBaseHandler(logger slog.Logger) *BaseHandler {
	return &BaseHandler{logger: logger}
}

type Handler struct {
	RegisterHandler *RegisterHandler
	LoginHandler    *LoginHandler
}

func NewHandler(services *service.Service, baseHandler BaseHandler) *Handler {
	return &Handler{
		RegisterHandler: NewRegisterHandler(services.RegisterService, baseHandler),
		LoginHandler:    NewLoginHandler(services.LoginService, baseHandler),
	}
}

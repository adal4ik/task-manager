package handlers

import (
	"net/http"
	"task-manager/internal/core/interfaces/driver"
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

}

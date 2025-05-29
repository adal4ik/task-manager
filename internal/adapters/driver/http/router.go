package web

import (
	"task-manager/internal/adapters/driver/http/handlers"

	"github.com/go-chi/chi/v5"
)

func Router(handler handlers.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/register", handler.RegisterHandler.RegisterUser)
	return r
}

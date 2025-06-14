package web

import (
	"task-manager/internal/adapters/driver/http/handlers"
	"task-manager/internal/adapters/driver/http/middleware"

	"github.com/go-chi/chi/v5"
)

func Router(handler handlers.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/register", handler.RegisterHandler.RegisterUser)
	r.Post("/login", handler.LoginHandler.LoginUser)
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthenticateJWT)
		r.Post("/task", handler.TaskHandler.CreateTask)
		r.Get("/tasks", handler.TaskHandler.GetTasks)
		r.Patch("/task/{task_id}", handler.TaskHandler.PatchTask)
		r.Delete("/task/{task_id}", handler.TaskHandler.DeleteTask)
	})
	return r
}

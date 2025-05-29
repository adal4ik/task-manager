package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"task-manager/internal/adapters/driven/database"
	"task-manager/internal/adapters/driven/database/repository"
	h "task-manager/internal/adapters/driver/http"
	"task-manager/internal/adapters/driver/http/handlers"
	"task-manager/internal/config"
	"task-manager/internal/core/service"
	"task-manager/internal/utils"
	"time"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()
	db := database.ConnectDB(ctx, cfg.Database)
	defer db.Close()
	logger, logFile := utils.Logger()
	defer logFile.Close()
	baseHandler := handlers.NewBaseHandler(*logger)
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handlers.NewHandler(services, *baseHandler)
	mux := h.Router(*handlers)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		log.Println("Server is running on port: http://localhost" + httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %s", err)
		}
	}()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
}

package database_test

import (
	"context"
	"task-manager/internal/adapters/driven/database"
	"task-manager/internal/config"
	"testing"
	"time"
)

func TestConnectDB(t *testing.T) {
	ctx := context.Background()
	cfg := config.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Name:     "taskmanager",
	}

	done := make(chan struct{})
	go func() {
		db := database.ConnectDB(ctx, cfg)
		if db == nil {
			t.Error("Expected non-nil db")
		}
		db.Close()
		close(done)
	}()

	select {
	case <-done:
		t.Log("ConnectDB completed successfully")
	case <-time.After(20 * time.Second):
		t.Fatal("ConnectDB timed out")
	}
}

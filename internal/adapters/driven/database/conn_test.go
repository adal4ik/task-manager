package database

import (
	"context"
	"os"
	"testing"
	"time"

	"task-manager/internal/config"
)

func TestConnectDB(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "taskmanager")

	cfg := config.Load()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := ConnectDB(ctx, cfg.Database)
	if db == nil {
		t.Fatal("db is nil")
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("failed to ping database: %v", err)
	}
}

package main

import (
	"context"
	"task-manager/internal/adapters/driven/database"
	"task-manager/internal/config"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()
	db := database.ConnectDB(ctx, cfg.Database)
	defer db.Close()
}

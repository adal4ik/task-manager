package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"task-manager/internal/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for PostgreSQL
)

func ConnectDB(ctx context.Context, cfg config.DatabaseConfig) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Fatalf("Timeout: failed to connect to DB within the deadline: %v", ctx.Err())
		case <-ticker.C:
			if err := db.PingContext(ctx); err == nil {
				log.Println("Successfully connected to the database!")
				return db
			} else {
				log.Println("Database not ready yet, retrying...")
			}
		}
	}
}

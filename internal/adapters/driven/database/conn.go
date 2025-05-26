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
	time.Sleep(5 * time.Second)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Successfully connected to the database!")
	return db
}

package config

import (
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     mustEnv("DB_HOST"),
			Port:     mustEnv("DB_PORT"),
			User:     mustEnv("DB_USER"),
			Password: mustEnv("DB_PASSWORD"),
			Name:     mustEnv("DB_NAME"),
		},
	}
}

func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("ENV variable %s is required but not set", key)
	}
	return val
}

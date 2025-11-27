// Package config provides configuration management for the application.
// It handles loading environment variables and setting up application configuration.
package config

import (
	"fmt"
	"log/slog"
	"os"
)

// ServerPort defines the default port for the HTTP server.
const ServerPort = "8080"

// DBDSN holds the database connection string constructed from environment variables.
var DBDSN string

// Load initializes the application configuration by reading environment variables
// and constructing the database connection string. Uses default values for missing variables.
func Load() error {
	dbHost := getEnv("DB_HOST", "db")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "qna_db")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	DBDSN = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
	)

	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		slog.Warn(fmt.Sprintf("environment variable %s is missing, setting default value: %s", key, defaultValue))
		return defaultValue
	}
	return value
}

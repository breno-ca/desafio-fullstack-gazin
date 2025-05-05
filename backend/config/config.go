package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	CorsAllowOrigin string
	MySQL           mySQLConfig
}

const (
	// Default Server Configurations
	SERVER_PORT       = ":8080"
	CORS_ALLOW_ORIGIN = "http://localhost:4200"
)

// NewConfig create a new config instance
func NewConfig() Config {
	// Optionally load variables from .env file if it exists
	godotenv.Load(".env")

	return Config{
		ServerPort:      getEnvOrDefault("SERVER_PORT", SERVER_PORT),
		CorsAllowOrigin: getEnvOrDefault("CORS_ALLOW_ORIGIN", CORS_ALLOW_ORIGIN),
	}
}

// getEnvOrDefault get values from environment variables or fallback to default
func getEnvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

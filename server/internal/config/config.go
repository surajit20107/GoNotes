package config

import (
  "os"
  "github.com/joho/godotenv"
  "github.com/surajit/notes-api/internal/logger"
)

type Config struct {
  DB_URL     string
  JWT_SECRET string
}

func LoadConfig() *Config {
  err := godotenv.Load()
  if err != nil {
    logger.Log.Warn("Warning: Failed to load .env file using system environment variables")
  }
  return &Config {
    DB_URL: getEnv("DB_URL"),
    JWT_SECRET: getEnv("JWT_SECRET"),
  }
}

// helper function to get environment variables
func getEnv(key string) string {
  value := os.Getenv(key)
  if value == "" {
    logger.Log.Error("Missing required environment variable:", key)
    os.Exit(1)
  }
  return value
}

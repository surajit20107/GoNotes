package config

import (
  "log"
  "os"
  "fmt"
  "github.com/joho/godotenv"
)

type Config struct {
  DB_URL     string
  JWT_SECRET string
}

func LoadConfig() *Config {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Warning: Failed to load .env file using system environment variables")
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
    log.Fatalf("Missing required environment variable: %s", key)
  }
  return value
}

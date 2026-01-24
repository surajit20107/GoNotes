package database

import (
  "os"
  "time"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "github.com/surajit/notes-api/internal/config"
  "github.com/surajit/notes-api/internal/logger"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) *gorm.DB {
  database, err := gorm.Open(postgres.Open(cfg.DB_URL), &gorm.Config{})
  if err != nil {
    logger.Log.Error("Database connection failed:", err)
    os.Exit(1)
  }

  // get underlying sql.DB for pool settings
  sqlDB, err := database.DB()
  if err != nil {
    logger.Log.Error("Failed to get database instance:", err)
    os.Exit(1)
  }

  // connection pool settings
  sqlDB.SetMaxIdleConns(10)
  sqlDB.SetMaxOpenConns(100)
  sqlDB.SetConnMaxLifetime(time.Hour)

  // optional: verify DB connection
  if err = sqlDB.Ping(); err != nil {
    logger.Log.Error("Failed to ping database", err)
  }

  logger.Log.Info("Database connected successfully ✅")

  // return the database instance
  DB = database
  return DB
}

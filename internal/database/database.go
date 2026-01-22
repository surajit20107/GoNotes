package database

import (
  "fmt"
  "log"
  "time"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "github.com/surajit/notes-api/internal/config"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) *gorm.DB {
  database, err := gorm.Open(postgres.Open(cfg.DB_URL), &gorm.Config{})
  if err != nil {
    log.Fatal("Database connection failed:", err)
  }

  // get underlying sql.DB for pool settings
  sqlDB, err := database.DB()
  if err != nil {
    log.Fatal("Failed to get database instance:", err)
  }

  // connection pool settings
  sqlDB.SetMaxIdleConns(10)
  sqlDB.SetMaxOpenConns(100)
  sqlDB.SetConnMaxLifetime(time.Hour)

  // optional: verify DB connection
  if err = sqlDB.Ping(); err != nil {
    log.Fatal("Failed to ping database", err)
  }

  fmt.Println("Database connected successfully ✅")

  // logger
  // &gorm.Config{
  //   Logger: logger.Default.LogMode(logger.Warn),
  // }

  // return the database instance
  DB = database
  return DB
}

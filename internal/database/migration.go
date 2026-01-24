package database

import (
  "os"
  "github.com/surajit/notes-api/internal/models"
  "github.com/surajit/notes-api/internal/logger"
  "gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
  err := db.AutoMigrate(
    &models.User{},
    &models.Note{},
  )
  
  if err != nil {
    logger.Log.Error("Failed to migrate Database:", err)
    os.Exit(1)
  }
  
}

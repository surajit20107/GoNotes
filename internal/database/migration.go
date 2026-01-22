package database

import (
  "log"
  "github.com/surajit/notes-api/internal/models"
  "gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
  err := db.AutoMigrate(
    &models.User{},
    &models.Note{},
  )
  
  if err != nil {
    log.Fatal("Failed to migrate Database:", err)
  }
  
}

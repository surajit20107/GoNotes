package models

import (
  "github.com/google/uuid"
  "time"
)

type Note struct {
  ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
  Title string `gorm:"not null" json:"title"`
  Content string `gorm:"not null" json:"content"`
  Author uuid.UUID `gorm:"not null" json:"author"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

type NoteDTO struct {
  Title string `json:"title" binding:"required,min=3"`
  Content string `json:"content" binding:"required,min=3"`
  Author uuid.UUID `json:"author"`
}

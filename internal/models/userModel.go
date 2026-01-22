package models

import (
  "github.com/google/uuid"
  "time"
)

type User struct {
  ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
  Name string `gorm:"not null" json:"name"`
  Email string `gorm:"unique;not null" json:"email"`
  Password string `gorm:"not null" json:"-"`
  Notes []Note `gorm:"foreignKey:Author" json:"notes"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//   u.ID = uuid.New()
//   return
// }

type SignupDTO struct {
  Name string `json:"name" binding:"required,min=3,max=20"`
  Email string `json:"email" binding:"required,email,min=15,max=50"`
  Password string `json:"password" binding:"required,min=8,max=20"`
}

type LoginDTO struct {
  Email string `json:"email" binding:"required,email,min=15,max=50"`
  Password string `json:"password" binding:"required,min=8,max=20"`
}

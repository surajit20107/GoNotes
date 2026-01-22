package services

import (
  "errors"
  "fmt"
  "github.com/surajit/notes-api/internal/database"
  "github.com/surajit/notes-api/internal/models"
  "github.com/surajit/notes-api/utils"
   "github.com/google/uuid"
)

type AuthService struct{}

func NewAuthService() *AuthService {
  return &AuthService{}
}

func (s *AuthService) RegisterUser(dto models.SignupDTO) (*models.User, error) {
  // check if the user already exists in the database
  var existingUser models.User
  if err := database.DB.Where("email = ?", dto.Email).First(&existingUser).Error; err == nil {
    return nil, errors.New("Email already registered")
  }

  // hash the password
  hashed, err := utils.HashPassword(dto.Password)
  if err != nil {
    return nil, err
  }

  // create a new user
  user := models.User{
    ID: uuid.New(),
    Name: dto.Name,
    Email: dto.Email,
    Password: hashed,
  }

  // save the user to the database
  if err := database.DB.Create(&user).Error; err != nil {
    fmt.Println("Error creating user: ", err.Error())
    return nil, err
  }
  return &user, nil
}

func (s *AuthService) LoginUser(dto models.LoginDTO) (*models.User, error) {
  var existingUser models.User

  // check if the user exists in the database
  if err := database.DB.Where("email = ?", dto.Email).First(&existingUser).Error; err != nil {
    return nil, errors.New("Invalid email or password")
  }

  // compare the password
  if err := utils.ComparePassword(existingUser.Password, dto.Password); err != nil {
    return nil, errors.New("Invalid email or password")
  }

  return &existingUser, nil
}

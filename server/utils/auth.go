package utils

import (
  "golang.org/x/crypto/bcrypt"
)

// convert user input plain text password to hashed password
func HashPassword(pass string) (string, error) {
  hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  }
  return string(hash), nil
}

// compares hashed password with user input plain text password
func ComparePassword(hash, pass string) error {
  return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}

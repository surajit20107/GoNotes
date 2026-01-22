package utils

import (
  "errors"
  "time"
  "github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
  UserID string `json:"user_id"`
  jwt.RegisteredClaims
}

func GenerateToken(userId, secret string) (string, error) {
  if secret == "" {
    return "", errors.New("JWT secret is required")
  }
  
  claims := JWTClaims{
    UserID: userId,
    RegisteredClaims: jwt.RegisteredClaims{
      Issuer: "notes-api",
      IssuedAt: jwt.NewNumericDate(time.Now()),
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString, secret string) (*JWTClaims, error) {
  if secret == "" {
    return nil, errors.New("JWT secret is required")
  }

  token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.New("unexpected signing method")
    }
    return []byte(secret), nil
  })

  if err != nil {
    return nil, err
  }

  claims, ok := token.Claims.(*JWTClaims)
  if !ok || !token.Valid {
    return nil, errors.New("invalid token")
  }

  return claims, nil
}

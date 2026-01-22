package middleware

import (
  "net/http"
  "strings"
  "github.com/gin-gonic/gin"
  "github.com/surajit/notes-api/utils"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
  return func(c *gin.Context) {
    var tokenString string
    authHeader := c.GetHeader("Authorization")
    if strings.HasPrefix(authHeader, "Bearer ") {
      tokenString = strings.TrimPrefix(authHeader, "Bearer ")
    }
    
    if tokenString == "" {
      cookie, err := c.Cookie("access_token")
      if err == nil {
      tokenString = cookie
      }
    }

    if tokenString == "" {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
          "error": "Authentication required",
      })
      return
    }
    
    claims, err := utils.VerifyToken(tokenString, secret)
    if err != nil {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid or expired token",
      })
      return
    }
    
    c.Set("user_id", claims.UserID)
    c.Next()
  }
}

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

    // extract token from header
    authHeader := c.GetHeader("Authorization")
    
    if authHeader != "" {
      if !strings.HasPrefix(authHeader, "Bearer ") {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
          "error": "Invalid authorization header format",
        })
        return
      }
      
      tokenString = strings.TrimPrefix(authHeader, "Bearer ")
    } else {
      // extract token from cookie
      cookie, err := c.Cookie("access_token")
      if err != nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
          "error": "Authentication required",
        })
        return
      }

      tokenString = cookie
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

    // set user id in context
    c.Set("user_id", claims.UserID)
    c.Next()
  }
}

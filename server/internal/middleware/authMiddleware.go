package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/surajit/notes-api/utils"
)

func addCORSHeaders(c *gin.Context) {
    origin := c.Request.Header.Get("Origin")
    if origin != "" {
        c.Header("Access-Control-Allow-Origin", origin)
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    }
}

func AuthMiddleware(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {

        var tokenString string

        // extract token from header
        authHeader := c.GetHeader("Authorization")

        if authHeader != "" {
            if !strings.HasPrefix(authHeader, "Bearer ") {
                addCORSHeaders(c)
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
                addCORSHeaders(c)
                c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                    "error": "Authentication required",
                })
                return
            }

            tokenString = cookie
        }

        if tokenString == "" {
            addCORSHeaders(c)
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Authentication required",
            })
            return
        }

        // validate token
        claims, err := utils.VerifyToken(tokenString, secret)
        if err != nil {
            addCORSHeaders(c)
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

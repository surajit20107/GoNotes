package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/surajit/notes-api/internal/controllers"
)

func AuthRoutes(r *gin.Engine) {
  auth := r.Group("/api/v1/auth")
  authController := controllers.NewAuthController()
  {
    auth.POST("/register", authController.Register)
    auth.POST("/login", authController.Login)
  }
}

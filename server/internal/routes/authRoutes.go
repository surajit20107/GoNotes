package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/surajit/notes-api/internal/controllers"
  "github.com/surajit/notes-api/internal/middleware"
  "github.com/surajit/notes-api/internal/config"
)

func AuthRoutes(r *gin.Engine, cfg *config.Config) {
  auth := r.Group("/api/v1/auth")
  authController := controllers.NewAuthController()
  {
    auth.POST("/register", authController.Register)
    auth.POST("/login", authController.Login)
    auth.POST("/logout", middleware.AuthMiddleware(cfg.JWT_SECRET), authController.Logout)
  }
}

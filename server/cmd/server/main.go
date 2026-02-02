package main

import (
        "time"
        
        "github.com/gin-contrib/cors"
        "github.com/gin-gonic/gin"
        "github.com/surajit/notes-api/internal/config"
        "github.com/surajit/notes-api/internal/database"
        "github.com/surajit/notes-api/internal/logger"
        "github.com/surajit/notes-api/internal/routes"
)

func main() {
  logger.Init()
  logger.Log.Info("Starting server...")
  
  cfg := config.LoadConfig()
  database.ConnectDB(cfg)
  database.AutoMigrate(database.DB)
  
  r := gin.Default()
  r.SetTrustedProxies(nil)
  frontend := "https://gonextpad.vercel.app"
  
  // middlewares
  r.Use(cors.New(cors.Config{
    AllowOrigins: []string{frontend},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
    AllowHeaders:     []string{"Content-Type", "Authorization", "Cookie"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
  }))
  
  // Routes
  r.GET("/", healthCheck)
  r.HEAD("/health", healthCheck)
  routes.AuthRoutes(r, cfg)
  routes.NoteRoutes(r, cfg)
  r.Run()
}

// health check method
func healthCheck(c *gin.Context) {
  c.JSON(200, gin.H{
    "success": true,
    "message": "Server up and running...🚀",
  })
}

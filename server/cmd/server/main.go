package main

import (
  "github.com/gin-gonic/gin"
  "github.com/surajit/notes-api/internal/config"
  "github.com/surajit/notes-api/internal/database"
  "github.com/surajit/notes-api/internal/routes"
  "github.com/surajit/notes-api/internal/logger"
)

func main() {
  logger.Init()
  logger.Log.Info("Starting server...")
  cfg := config.LoadConfig()
  database.ConnectDB(cfg)
  database.AutoMigrate(database.DB)
  r := gin.Default()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())
  r.GET("/", healthCheck)
  r.HEAD("/health", healthCheck)
  routes.AuthRoutes(r, cfg)
  routes.NoteRoutes(r, cfg)
  r.Run()
}

func healthCheck(c *gin.Context) {
  c.JSON(200, gin.H{
    "success": true,
    "message": "Server up and running...🚀",
  })
}

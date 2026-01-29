package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "github.com/surajit/notes-api/internal/config"
  "github.com/surajit/notes-api/internal/database"
  "github.com/surajit/notes-api/internal/routes"
  "github.com/surajit/notes-api/internal/logger"
  "time"
)

func main() {
  logger.Init()
  logger.Log.Info("Starting server...")
  
  cfg := config.LoadConfig()
  database.ConnectDB(cfg)
  database.AutoMigrate(database.DB)
  
  r := gin.Default()

  // temporary middleware
  r.Use(func(c *gin.Context) {
    origin := c.Request.Header.Get("Origin")
    if origin != "" {
        logger.Log.Info("Request Origin:", origin)
    }
    c.Next()
})
  
  // middlewares
  r.Use(cors.New(cors.Config{
    AllowOrigins: []string{
      "https://go-notes-frontend.vercel.app",
      "https://c9e6293d-a3f2-4638-a088-507983afa80f-00-1z3rnias14xhz.sisko.replit.dev",
    },
    AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
    AllowHeaders: []string{ "Content-Type", "Authorization"},
    ExposeHeaders: []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
  }))

  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  // r.OPTIONS("/*path", func(c *gin.Context) {
  //   c.AbortWithStatus(204)
  // })
  
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

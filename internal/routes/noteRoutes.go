package routes

import (
  "github.com/gin-gonic/gin"
)

func NoteRoutes(r *gin.Engine) {
  note := r.Group("/api/v1/notes")
  {
    note.GET("/", test)
    // note.GET("/:id", test)
    // note.POST("/", test)
    // note.PUT("/:id", test)
    // note.DELETE("/:id", test)
  }
}

func test(c *gin.Context) {
  c.JSON(200, gin.H{
    "success": true,
    "message": "API is working...⚙️",
  })
}

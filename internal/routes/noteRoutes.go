package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/surajit/notes-api/internal/controllers"
  "github.com/surajit/notes-api/internal/middleware"
  "github.com/surajit/notes-api/internal/config"
  "github.com/surajit/notes-api/internal/services"
)

func NoteRoutes(r *gin.Engine, cfg *config.Config) {
  note := r.Group("/api/v1/notes")
  // apply auth middleware to all routes in note group
  note.Use(middleware.AuthMiddleware(cfg.JWT_SECRET))
  noteService := services.NewNoteService()
  noteController := controllers.NewNoteController(noteService)
  {
    note.GET("/", noteController.GetNotes)
    note.GET("/:id", noteController.GetNoteById)
    note.POST("/", noteController.CreateNote)
    note.PUT("/:id", noteController.UpdateNote)
    note.DELETE("/:id", noteController.DeleteNote)
  }
}

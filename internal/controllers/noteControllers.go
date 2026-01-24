package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/surajit/notes-api/internal/models"
  "github.com/surajit/notes-api/internal/services"
  "github.com/google/uuid"
)

type NoteController struct {
  noteService *services.NoteService
}

func NewNoteController(noteService *services.NoteService) *NoteController {
  return &NoteController {
    noteService: noteService,
  }
}

// get all notes for a user
func (nc *NoteController) GetNotes(c *gin.Context) {
  // get current logged in user id using middleware
  userId := c.GetString("user_id")
  if userId == "" {
    c.JSON(http.StatusUnauthorized, gin.H{
      "success": false,
      "error": "Unauthorized",
    })
    return
  }

  // parse user id to uuid
  uid, err := uuid.Parse(userId)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": "Invalid user id, please login again",
    })
    return
  }

  // call service layer to get all notes by user id
  notes, err := nc.noteService.GetNotes(uid)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "success": false,
      "error": "Failed to fetch notes",
    })
    return
  }

  // return success response with notes
  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "notes": notes,
  })
}

// create a new note
func (nc *NoteController) CreateNote(c *gin.Context) {
  // get current logged in user id using middleware
  userId := c.GetString("user_id")
  if userId == "" {
    c.JSON(http.StatusUnauthorized, gin.H{
      "success": false,
      "error": "Unauthorized",
    })
    return
  }

  // parse user id to uuid
  uid, err := uuid.Parse(userId)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": "Invalid user id, please login again",
    })
    return
  }
  
  var noteDto models.NoteDTO

  if err := c.ShouldBindJSON(&noteDto); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  note, err := nc.noteService.CreateNote(noteDto, uid)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusCreated, gin.H{
    "success": true,
    "message": "Note created successfully",
    "note": note,
  })
}

// get notes by id
func (nc *NoteController) GetNoteById(c *gin.Context) {
  userId := c.GetString("user_id")
  if userId == "" {
    c.JSON(http.StatusUnauthorized, gin.H{
      "success": false,
      "error": "Unauthorized",
    })
    return
  }

  uid, err := uuid.Parse(userId)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": "Invalid user id, please login again",
    })
    return
  }

  noteId := c.Param("id")
  note, err := nc.noteService.GetNoteById(uid, noteId)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "note": note,
  })
}

// update note by id
func (nc *NoteController) UpdateNote(c *gin.Context) {
  userId := c.GetString("user_id")
  if userId == "" {
    c.JSON(http.StatusUnauthorized, gin.H{
      "success": false,
      "error": "Unauthorized",
    })
    return
  }

  uid, err := uuid.Parse(userId)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": "Invalid user id, please login again",
    })
    return
  }

  noteId := c.Param("id")
  var updatedData map[string]interface{}
  if err := c.ShouldBindJSON(&updatedData); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  note, err := nc.noteService.UpdateNote(uid, noteId, &updatedData)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "message": "Note updated successfully",
    "note": note,
  })
}

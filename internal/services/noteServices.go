package services

import (
  "github.com/surajit/notes-api/internal/models"
  "github.com/surajit/notes-api/internal/database"
  "github.com/google/uuid"
)

type NoteService struct {}

func NewNoteService() *NoteService {
  return &NoteService{}
}

// fetch all notes for a user from database and return them
func (ns *NoteService) GetNotes(userId uuid.UUID) ([]models.Note, error) {
  // declare a slice of notes
  var notes []models.Note

  // find all notes by user id
  err := database.DB.Where("author = ?", userId).Find(&notes).Error
  if err != nil {
    return nil, err
  }

  // return notes if no error
  return notes, nil
}

// create a new note to database and return the created note
func (ns *NoteService) CreateNote(noteDto models.NoteDTO, userId uuid.UUID) (models.Note, error) {
  var note models.Note
  note = models.Note{
    ID: uuid.New(),
    Title: noteDto.Title,
    Content: noteDto.Content,
    Author: userId,
  }
  err := database.DB.Create(&note).Error
  if err != nil {
    return models.Note{}, err
  }

  return note, nil
}

// fetch and return a note by id
func (ns *NoteService) GetNoteById(userId uuid.UUID, noteId string) (models.Note, error) {
  var note models.Note
  err := database.DB.Where("id = ? AND author = ?", noteId, userId).First(&note).Error
  if err != nil {
    return models.Note{}, err
  }

  return note, nil
}

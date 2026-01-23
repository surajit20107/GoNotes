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

// create a new note
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

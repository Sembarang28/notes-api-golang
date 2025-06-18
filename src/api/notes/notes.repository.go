package notes

import (
	"notes-management-api/src/api/notes/dto"
	"notes-management-api/src/models"
)

type NotesRepository interface {
	Create(notes *models.Notes) error
	ReadAll(search, userId string, tags []string) ([]dto.NotesResponse, error)
	ReadOne(id, userId string) (*dto.NotesResponse, error)
	Update(notes *models.Notes) error
	Delete(id, userId string) error
}

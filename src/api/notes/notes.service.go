package notes

import "notes-management-api/src/api/notes/dto"

type NoteService interface {
	Create(request *dto.NotesRequest, userId string) error
	ReadAll(search, userId string, tags []string) ([]dto.NotesResponse, error)
	ReadOne(id, userId string) (*dto.NotesResponse, error)
	Update(request *dto.NotesRequest, id, userId string) error
	Delete(id, userId string) error
}

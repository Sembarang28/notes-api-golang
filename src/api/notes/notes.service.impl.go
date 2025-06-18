package notes

import (
	"fmt"
	"notes-management-api/src/api/notes/dto"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"

	"github.com/go-playground/validator/v10"
)

type NoteServiceImpl struct {
	validator      *validator.Validate
	noteRepository NotesRepository
}

func NewNoteService(validator *validator.Validate, noteRepository NotesRepository) NoteService {
	return &NoteServiceImpl{
		validator:      validator,
		noteRepository: noteRepository,
	}
}

func (s NoteServiceImpl) Create(request *dto.NotesRequest, userId string) error {
	if err := s.validator.Struct(request); err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	notes := &models.Notes{
		Name:       request.Name,
		Notes:      request.Notes,
		CategoryID: request.CategoryID,
		UserID:     userId,
		Tags:       request.Tags,
	}

	return s.noteRepository.Create(notes)
}

func (s NoteServiceImpl) ReadAll(search, userId string, tags []string) ([]dto.NotesResponse, error) {
	notes, err := s.noteRepository.ReadAll(search, userId, tags)

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s NoteServiceImpl) ReadOne(id, userId string) (*dto.NotesResponse, error) {
	note, err := s.noteRepository.ReadOne(id, userId)

	if err != nil {
		return nil, err
	}

	return note, nil
}

func (s NoteServiceImpl) Update(request *dto.NotesRequest, id, userId string) error {
	if err := s.validator.Struct(request); err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	notes := &models.Notes{
		ID:         id,
		Name:       request.Name,
		Notes:      request.Notes,
		CategoryID: request.CategoryID,
		UserID:     userId,
		Tags:       request.Tags,
	}

	return s.noteRepository.Update(notes)
}

func (s NoteServiceImpl) Delete(id, userId string) error {
	return s.noteRepository.Delete(id, userId)
}

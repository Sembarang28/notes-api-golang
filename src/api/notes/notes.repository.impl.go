package notes

import (
	"encoding/json"
	"fmt"
	"notes-management-api/src/api/notes/dto"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"

	"gorm.io/gorm"
)

type NotesRepositoryImpl struct {
	db *gorm.DB
}

func NewNotesRepository(db *gorm.DB) NotesRepository {
	return &NotesRepositoryImpl{db: db}
}

func (r NotesRepositoryImpl) Create(notes *models.Notes) error {
	err := r.db.Create(&notes).Error

	if err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return nil
}

func (r NotesRepositoryImpl) ReadAll(search, userId string, tags []string) ([]dto.NotesResponse, error) {
	var rawNotes []dto.NotesResponse

	query := r.db.Model(&models.Notes{}).
		Select(`notes.id::text, notes.name, notes.notes, notes.category_id, category.name as category_name, notes.tags::text`).
		Joins("JOIN category ON category.id = notes.category_id").
		Where("notes.user_id = ?", userId)

	// Search by name
	if search != "" {
		query = query.Where("notes.name ILIKE ?", "%"+search+"%")
	}

	// Filter by tags if provided
	if len(tags) > 0 {
		tagJson, err := json.Marshal(tags)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}
		query = query.Where("notes.tags @> ?::jsonb", string(tagJson))
	}

	// Execute query into rawNotes
	result := query.Scan(&rawNotes)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, helpers.ErrNotFound
	}

	// Decode raw JSON tags
	notes := make([]dto.NotesResponse, 0, len(rawNotes))
	for _, n := range rawNotes {
		var decodedTags []string
		if err := json.Unmarshal([]byte(n.RawTags), &decodedTags); err != nil {
			decodedTags = []string{} // fallback to empty
		}
		n.Tags = decodedTags
		notes = append(notes, n)
	}

	return notes, nil
}

func (r NotesRepositoryImpl) ReadOne(id, userId string) (*dto.NotesResponse, error) {
	var note dto.NotesResponse

	err := r.db.Model(&models.Notes{}).
		Select("notes.id::text, notes.name, notes.notes, notes.category_id, category.name as category_name, notes.tags").
		Joins("JOIN category ON category.id = notes.category_id").
		Where("notes.id = ? AND notes.user_id = ?", id, userId).
		Scan(&note).Error

	if err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	if note.ID == "" {
		return nil, fmt.Errorf("%w: note with id '%s' not found", helpers.ErrNotFound, id)
	}

	var decodedTags []string
	if err := json.Unmarshal([]byte(note.RawTags), &decodedTags); err != nil {
		decodedTags = []string{} // fallback to empty
	}
	note.Tags = decodedTags

	return &note, nil
}

func (r NotesRepositoryImpl) Update(notes *models.Notes) error {
	result := r.db.Model(&models.Notes{}).
		Where("id = ? AND user_id = ?", notes.ID, notes.UserID).
		Updates(map[string]interface{}{
			"name":        notes.Name,
			"notes":       notes.Notes,
			"category_id": notes.CategoryID,
			"tags":        notes.Tags,
		})

	if result.Error != nil {
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: note with id '%s' not found", helpers.ErrNotFound, notes.ID)
	}

	return nil
}

func (r NotesRepositoryImpl) Delete(id, userId string) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Notes{})

	if result.Error != nil {
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: note with id '%s' not found", helpers.ErrNotFound, id)
	}

	return nil
}

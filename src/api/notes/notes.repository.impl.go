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

func NewCategoryRepository(db *gorm.DB) NotesRepository {
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
	var notes []dto.NotesResponse

	query := r.db.Model(&models.Notes{}).
		Select("notes.id::text, notes.name, notes.notes, notes.category_id, categories.name as category_name, notes.tags").
		Joins("JOIN categories ON categories.id = notes.category_id").
		Where("notes.user_id = ?", userId)

	// search by name
	if search != "" {
		query = query.Where("notes.name ILIKE ?", "%"+search+"%")
	}

	// filter tags if provided
	if len(tags) > 0 {
		// Convert to PostgreSQL array syntax: '{"tag1","tag2"}'
		tagQuery := "notes.tags @> ?::jsonb"
		tagJson, _ := json.Marshal(tags) // convert to `["tag1","tag2"]`
		query = query.Where(tagQuery, string(tagJson))
	}

	if err := query.Scan(&notes).Error; err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}

	return notes, nil
}

func (r NotesRepositoryImpl) ReadOne(id, userId string) (*dto.NotesResponse, error) {
	var note dto.NotesResponse

	err := r.db.Model(&models.Notes{}).
		Select("notes.id::text, notes.name, notes.notes, notes.category_id, categories.name as category_name, notes.tags").
		Joins("JOIN categories ON categories.id = notes.category_id").
		Where("notes.id = ? AND notes.user_id = ?", id, userId).
		Scan(&note).Error

	if err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	if note.ID == "" {
		return nil, fmt.Errorf("%w: note with id '%s' not found", helpers.ErrNotFound, id)
	}

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

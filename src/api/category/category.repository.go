package category

import (
	"notes-management-api/src/api/category/dto"
	"notes-management-api/src/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	ReadAll(search string) ([]dto.CategoryResponse, error)
	Read(id string) (*dto.CategoryResponse, error)
	Update(category *models.Category) error
	Delete(id string) error
}

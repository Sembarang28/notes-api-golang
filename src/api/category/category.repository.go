package category

import "notes-management-api/src/models"

type CategoryRepository interface {
	Create(category *models.Category) error
	ReadAll(search string) ([]models.Category, error)
	Read(id string) (models.Category, error)
	Update(category *models.Category) error
	Delete(id string) error
}

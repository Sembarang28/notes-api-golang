package category

import (
	"fmt"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"
	"strings"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (r CategoryRepositoryImpl) Create(category *models.Category) error {
	if err := r.db.Create(category).Error; err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return nil
}

func (r CategoryRepositoryImpl) ReadAll(search string) ([]models.Category, error) {
	var categories []models.Category
	if search != "" {
		err := r.db.Where("name ILIKE ?", "%"+search+"%").Find(&categories).Error
		if err != nil {
			return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}
	} else {
		err := r.db.Find(&categories).Error
		if err != nil {
			return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}
	}
	return categories, nil
}

func (r CategoryRepositoryImpl) Read(id string) (models.Category, error) {
	category := models.Category{}
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return category, fmt.Errorf("%w: category with id '%s' not found", helpers.ErrNotFound, id)
		}
		return category, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return category, nil
}

func (r CategoryRepositoryImpl) Update(category *models.Category) error {
	if err := r.db.Where("id = ?", category.ID).Save(category).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return fmt.Errorf("'%w': category with id '%s' not found", helpers.ErrNotFound, category.ID)
		}
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return nil
}

func (r CategoryRepositoryImpl) Delete(id string) error {
	if err := r.db.Delete(&models.Category{}, "id = ?", id).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return fmt.Errorf("'%w': category with id '%s' not found", helpers.ErrNotFound, id)
		}
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return nil
}

package category

import (
	"fmt"
	"notes-management-api/src/api/category/dto"
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

func (r CategoryRepositoryImpl) ReadAll(search string) ([]dto.CategoryResponse, error) {
	var categories []dto.CategoryResponse
	if search != "" {
		err := r.db.Model(&models.Category{}).
			Select("id::text, name, description").
			Where("name ILIKE ?", "%"+search+"%").
			Scan(&categories).Error
		if err != nil {
			return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}
	} else {
		err := r.db.Model(&models.Category{}).
			Select("id::text, name, description").
			Scan(&categories).Error
		if err != nil {
			return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}
	}
	return categories, nil
}

func (r CategoryRepositoryImpl) Read(id string) (*dto.CategoryResponse, error) {
	category := dto.CategoryResponse{}
	if err := r.db.Model(&models.Category{}).
		Select("id::text, name, description").
		Where("id = ?", id).
		Scan(&category).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return &category, fmt.Errorf("%w: category with id '%s' not found", helpers.ErrNotFound, id)
		}
		return &category, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return &category, nil
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

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

func (r CategoryRepositoryImpl) ReadAll(search, userId string) ([]dto.CategoryResponse, error) {
	var categories []dto.CategoryResponse

	query := r.db.Model(&models.Category{}).
		Select("id::text, name, description").
		Where("user_id = ?", userId) // Always filter by userId

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	err := query.Scan(&categories).Error
	if err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}

	return categories, nil
}

func (r CategoryRepositoryImpl) Read(id, userId string) (*dto.CategoryResponse, error) {
	category := dto.CategoryResponse{}

	err := r.db.Model(&models.Category{}).
		Select("id::text, name, description").
		Where("id = ? AND user_id = ?", id, userId).
		Scan(&category).Error

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return &category, fmt.Errorf("%w: category with id '%s' not found", helpers.ErrNotFound, id)
		}
		return &category, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}

	return &category, nil
}

func (r CategoryRepositoryImpl) Update(category *models.Category) error {
	if err := r.db.Where("id = ? AND user_id = ?", category.ID, category.UserID).Save(category).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return fmt.Errorf("'%w': category with id '%s' not found", helpers.ErrNotFound, category.ID)
		}
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return nil
}

func (r CategoryRepositoryImpl) Delete(id, userId string) error {
	if err := r.db.Delete(&models.Category{}, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return fmt.Errorf("'%w': category with id '%s' not found", helpers.ErrNotFound, id)
		}
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return nil
}

package category

import (
	"fmt"
	"notes-management-api/src/api/category/dto"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	categoryRepository CategoryRepository
	validator          *validator.Validate
}

func NewCategoryService(categoryRepository CategoryRepository, validator *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
		validator:          validator,
	}
}

func (s CategoryServiceImpl) Create(request *dto.CategoryRequest, userId string) error {
	if err := s.validator.Struct(request); err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	category := models.Category{
		Name:        request.Name,
		UserID:      userId,
		Description: request.Description,
	}

	return s.categoryRepository.Create(&category)
}

func (s CategoryServiceImpl) ReadAll(search, userId string) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryRepository.ReadAll(search, userId)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s CategoryServiceImpl) Read(id, userId string) (*dto.CategoryResponse, error) {
	category, err := s.categoryRepository.Read(id, userId)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s CategoryServiceImpl) Update(request *dto.CategoryRequest, id, userId string) error {
	if err := s.validator.Struct(request); err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	category := models.Category{
		ID:          id,
		UserID:      userId,
		Name:        request.Name,
		Description: request.Description,
	}

	return s.categoryRepository.Update(&category)
}

func (s CategoryServiceImpl) Delete(id, userId string) error {
	return s.categoryRepository.Delete(id, userId)
}

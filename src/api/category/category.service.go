package category

import "notes-management-api/src/api/category/dto"

type CategoryService interface {
	Create(request *dto.CategoryRequest, userId string) error
	ReadAll(search string) ([]dto.CategoryResponse, error)
	Read(id string) (*dto.CategoryResponse, error)
	Update(request *dto.CategoryRequest, id, userId string) error
	Delete(id string) error
}

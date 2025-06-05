package dto

import "mime/multipart"

type UserUpdateRequest struct {
	Name  string                `form:"name" validate:"required"`
	Email string                `form:"email" validate:"required,email"`
	Photo *multipart.FileHeader `form:"photo" validate:"omitempty"`
}

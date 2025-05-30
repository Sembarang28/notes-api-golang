package dto

import "mime/multipart"

type UserUpdateRequest struct {
	Name            string                `json:"name" validate:"required"`
	Email           string                `json:"email" validate:"required,email"`
	Password        string                `json:"password" validate:"required,min=8"`
	ConfirmPassword string                `json:"confirmPassword" validate:"required,eqfield=Password"`
	Photo           *multipart.FileHeader `json:"image,omitempty" form:"photo" validate:"omitempty"`
}

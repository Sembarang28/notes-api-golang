package dto

type UserUpdateRequest struct {
	Name  string  `form:"name" validate:"required"`
	Email string  `form:"email" validate:"required,email"`
	Photo *string `form:"photo" validate:"omitempty,url"`
}

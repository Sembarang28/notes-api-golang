package dto

type UserUpdatePasswordRequest struct {
	OldPassword        string `json:"oldPassword" validate:"required,min=8"`
	NewPassword        string `json:"newPassword" validate:"required,min=8"`
	ConfirmNewPassword string `json:"confirmPassword" validate:"required,eqfield=NewPassword"`
}

package dto

type UserRefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

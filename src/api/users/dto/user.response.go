package dto

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhotoUrl string `json:"photo_url"`
}

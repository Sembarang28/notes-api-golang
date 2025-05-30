package dto

type UserProfileResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Photo string `json:"photo,omitempty"`
}

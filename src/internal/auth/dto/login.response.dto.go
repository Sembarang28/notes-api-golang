package dto

type UserLoginResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhotoUrl     string `json:"photo_url"`
}

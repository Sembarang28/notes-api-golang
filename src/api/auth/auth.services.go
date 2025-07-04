package auth

import "notes-management-api/src/api/auth/dto"

type AuthService interface {
	Register(request *dto.UserRegistrationRequest) error
	Login(request *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	RefreshToken(token string) (*dto.UserRefreshResponse, error)
	Logout(token string) error
}

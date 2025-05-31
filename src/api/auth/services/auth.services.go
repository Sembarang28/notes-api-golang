package services

import "notes-management-api/src/api/auth/dto"

type AuthService interface {
	Register(request *dto.UserRegistrationRequest) error
	Login(request *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
}

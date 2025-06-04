package services

import (
	"notes-management-api/src/api/users/dto"
)

type UserService interface {
	GetUserById(id string) (*dto.UserResponse, error)
	UpdateUser(id string, request *dto.UserUpdateRequest) error
	UpdateUserPassword(id string, request *dto.UserUpdatePasswordRequest) error
}

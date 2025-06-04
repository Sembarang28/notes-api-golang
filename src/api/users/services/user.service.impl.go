package services

import (
	"fmt"
	"notes-management-api/src/api/users/dto"
	"notes-management-api/src/api/users/repository"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	validator      *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		validator:      validator,
	}
}

func (s *UserServiceImpl) GetUserById(id string) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		Name:     user.Name,
		Email:    user.Email,
		PhotoUrl: *user.Photo,
	}, nil
}

func (s *UserServiceImpl) UpdateUser(id string, request *dto.UserUpdateRequest) error { // Ensure the user ID is set to the correct value
	if err := s.validator.Struct(request); err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	user := &models.User{
		ID:    id,
		Email: request.Email,
		Name:  request.Name,
	}

	return s.userRepository.Update(user)
}

func (s *UserServiceImpl) UpdateUserPassword(id string, request *dto.UserUpdatePasswordRequest) error {
	if err := s.validator.Struct(request); err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	data, err := s.userRepository.FindById(id)
	if err != nil {
		return err
	}

	match, err := helpers.CheckPasswordHash(request.OldPassword, data.Password)
	if err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}

	if !match {
		return fmt.Errorf("%w: old password is incorrect", helpers.ErrClientError)
	}

	// Hash the new password
	hashedPassword, err := helpers.HashPassword(request.NewPassword)
	if err != nil {
		return err
	}

	user := &models.User{
		ID:       id,
		Password: hashedPassword,
	}

	return s.userRepository.Update(user)
}

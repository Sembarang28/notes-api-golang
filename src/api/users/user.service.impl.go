package users

import (
	"fmt"
	"notes-management-api/src/api/users/dto"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	userRepository UserRepository
	validator      *validator.Validate
}

func NewUserService(userRepository UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		validator:      validator,
	}
}

func (s *UserServiceImpl) GetUserById(id string) (*dto.UserResponse, error) {
	fmt.Println(id)
	user, err := s.userRepository.FindById(id)
	fmt.Println(user.Name)

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

	var photoPath string
	if request.Photo != nil {
		src, err := request.Photo.Open()
		if err != nil {
			return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}
		defer src.Close()

		ext := strings.ToLower(filepath.Ext(request.Photo.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return fmt.Errorf("%w: unsupported file type %s", helpers.ErrClientError, ext)
		}

		img, err := imaging.Decode(src)
		if err != nil {
			return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}

		os.MkdirAll("public", os.ModePerm)
		filename := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102T150405.000"), uuid.New().String(), ext)
		savePath := filepath.Join("public", filename)
		if err := imaging.Save(img, savePath, imaging.JPEGQuality(80)); err != nil {
			return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}

		savePath = filepath.ToSlash(savePath) // Ensure the path is in the correct format for web access
	}

	user := &models.User{
		ID:    id,
		Email: request.Email,
		Name:  request.Name,
		Photo: &photoPath,
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

package users

import (
	"fmt"
	"notes-management-api/src/api/users/dto"
	"notes-management-api/src/helpers"
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

func (s *UserServiceImpl) UpdateUser(id string, request *dto.UserUpdateRequest) error {
	if err := s.validator.Struct(request); err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	user, err := s.userRepository.FindById(id)
	if err != nil {
		return err
	}

	var photoPath string

	// If new photo is uploaded
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

		// Delete previous photo if exists
		if user.Photo != nil && *user.Photo != "" {
			if err := os.Remove(*user.Photo); err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("%w: failed to delete previous photo: %v", helpers.ErrInternalServer, err)
			}
		}

		// Create public dir and generate filename
		os.MkdirAll("public", os.ModePerm)
		filename := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102T150405.000"), uuid.New().String(), ext)
		savePath := filepath.Join("public", filename)

		// Save compressed image
		if err := imaging.Save(img, savePath, imaging.JPEGQuality(80)); err != nil {
			return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}

		// Normalize path for web
		photoPath = filepath.ToSlash(savePath)
	} else {
		if user.Photo != nil {
			photoPath = *user.Photo // Keep old photo if no new one uploaded
		} else {
			photoPath = ""
		}
	}

	return s.userRepository.Update(id, request.Name, request.Email, photoPath)
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

	return s.userRepository.UpdatePassword(id, hashedPassword)
}

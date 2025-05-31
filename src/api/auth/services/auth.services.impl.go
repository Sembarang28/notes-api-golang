package services

import (
	"errors"
	"fmt"
	"notes-management-api/src/api/auth/dto"
	"notes-management-api/src/api/auth/repository"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
	validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		authRepository: authRepository,
		validate:       validate,
	}
}

func (s *AuthServiceImpl) Register(request *dto.UserRegistrationRequest) error {
	// Validate the request
	if err := s.validate.Struct(request); err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	// Hash the password
	hash, err := helpers.HashPassword(request.Password)
	if err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}

	// Save the user
	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hash,
	}

	return s.authRepository.Save(user)
}

func (s *AuthServiceImpl) Login(request *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	// Validate the request
	if err := s.validate.Struct(request); err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrValidation, err)
	}

	// Find user by email
	user, err := s.authRepository.FindByEmail(request.Email)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrNotFound):
			return nil, fmt.Errorf("%w: user with email '%s' not found", helpers.ErrNotFound, request.Email)
		default:
			return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
		}
	}

	// Verify password
	match, err := helpers.CheckPasswordHash(request.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	if !match {
		return nil, helpers.ErrUnauthorized
	}

	// Generate JWT Access Token
	accessToken, err := helpers.NewAccessToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}

	// Save session
	sessionID := uuid.New().String()

	refreshToken, err := helpers.NewRefreshToken(sessionID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}

	session := &models.Session{
		ID:        string(sessionID),
		UserID:    user.ID,
		Token:     refreshToken.Token,
		ExpiresAt: refreshToken.ExpiresAt,
		IssuedAt:  refreshToken.IssuedAt,
	}

	if err := s.authRepository.SaveSession(session); err != nil {
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}

	return &dto.UserLoginResponse{
		RefreshToken: refreshToken.Token,
		AccessToken:  accessToken,
		Name:         user.Name,
		Email:        user.Email,
		PhotoUrl:     *user.Photo,
	}, nil
}

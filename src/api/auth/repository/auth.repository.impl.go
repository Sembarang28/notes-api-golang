package repository

import (
	"fmt"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) Save(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			return fmt.Errorf("%w: email '%s' already exists", helpers.ErrEmailAlreadyExists, user.Email)
		}
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return nil
}

func (r *AuthRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("%w: user with email '%s' not found", helpers.ErrNotFound, email)
		}
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return &user, nil
}

func (r *AuthRepositoryImpl) SaveSession(session *models.Session) error {
	if err := r.db.Create(session).Error; err != nil {
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return nil
}

func (r *AuthRepositoryImpl) FindSessionByIDAndToken(sessionID, token string) (*models.Session, error) {
	var session models.Session
	if err := r.db.Where("id = ? AND token = ?", sessionID, token).First(&session).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("%w: session with id '%s' and token '%s' not found", helpers.ErrNotFound, sessionID, token)
		}
		return nil, fmt.Errorf("%w: %v", helpers.ErrInternalServer, err)
	}
	return &session, nil
}

func (r *AuthRepositoryImpl) UpdateSession(token string) error {
	result := r.db.Where("token = ?", token).Updates(&models.Session{Revoked: true})

	if result.Error != nil {
		return fmt.Errorf("%w: %v", helpers.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: session with token '%s' not found", helpers.ErrNotFound, token)
	}
	return nil
}

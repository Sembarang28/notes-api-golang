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

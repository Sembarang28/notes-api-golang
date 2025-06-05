package users

import (
	"fmt"
	"notes-management-api/src/helpers"
	"notes-management-api/src/models"
	"strings"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindById(id string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, fmt.Errorf("'%w': user with id '%s' not found", helpers.ErrNotFound, id)
		}
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Update(id, name, email, photo string) error {
	err := r.db.Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":  name,
			"email": email,
			"photo": photo,
		}).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return fmt.Errorf("'%w': user with id '%s' not found", helpers.ErrNotFound, id)
		}

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("'%w': email '%s' already exists", helpers.ErrEmailAlreadyExists, email)
		}
	}

	return nil
}

func (r *UserRepositoryImpl) UpdatePassword(id, newPassword string) error {
	err := r.db.Model(&models.User{}).
		Where("id = ?", id).
		Update("password", newPassword).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return fmt.Errorf("'%w': user with id '%s' not found", helpers.ErrNotFound, id)
		}
	}

	return nil
}

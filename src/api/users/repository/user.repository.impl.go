package repository

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
	err := r.db.First(&user, "id = ?", id).Error
	if strings.Contains(err.Error(), "record not found") {
		return nil, fmt.Errorf("'%w': user with id '%s' not found", helpers.ErrNotFound, id)
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	err := r.db.Where("id = ?", user.ID).Save(user).Error
	if strings.Contains(err.Error(), "record not found") {
		return fmt.Errorf("'%w': user with id '%s' not found", helpers.ErrNotFound, user.ID)
	}

	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return fmt.Errorf("'%w': email '%s' already exists", helpers.ErrEmailAlreadyExists, user.Email)
	}

	return nil
}

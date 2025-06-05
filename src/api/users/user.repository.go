package users

import (
	"notes-management-api/src/models"
)

type UserRepository interface {
	FindById(id string) (*models.User, error)
	Update(user *models.User) error
}

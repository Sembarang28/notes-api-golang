package users

import (
	"notes-management-api/src/models"
)

type UserRepository interface {
	FindById(id string) (*models.User, error)
	Update(id, name, email, photo string) error
	UpdatePassword(id, newPassword string) error
}

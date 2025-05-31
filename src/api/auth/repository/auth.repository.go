package repository

import "notes-management-api/src/models"

type AuthRepository interface {
	Save(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	SaveSession(session *models.Session) error
}

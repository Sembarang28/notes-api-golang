package auth

import "notes-management-api/src/models"

type AuthRepository interface {
	Save(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	SaveSession(session *models.Session) error
	FindSessionByIDAndToken(sessionID, token string) (*models.Session, error)
	UpdateSession(token string) error
}

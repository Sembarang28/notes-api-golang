package models

import (
	"time"
)

type Session struct {
	ID        string    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    string    `gorm:"column:user_id;not null"`
	Token     string    `gorm:"column:token;not null;uniqueIndex"`
	IssuedAt  time.Time `gorm:"column:issued_at;not null;autoCreateTime;<-:create"`
	ExpiresAt time.Time `gorm:"column:expires_at;not null"`
	Revoked   bool      `gorm:"column:revoked;not null;default:false"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
	User      User      `gorm:"foreignKey:user_id;references:id;OnDelete:CASCADE"`
}

func (s *Session) TableName() string {
	return "session"
}

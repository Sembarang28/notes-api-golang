package entity

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	ID        string `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    string `gorm:"column:user_id;not null"`
	token     string `gorm:"column:token;not null;uniqueIndex"`
	issuedAt  string `gorm:"column:issued_at;not null;autoCreateTime;<-:create"`
	ExpiresAt string `gorm:"column:expires_at;not null"`
	revoked   bool   `gorm:"column:revoked;not null;default:false"`
	CreatedAt string `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt string `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
	User      User   `gorm:"foreignKey:user_id;references:id;OnDelete:CASCADE"`
}

func (s *Session) TableName() string {
	return "sessions"
}

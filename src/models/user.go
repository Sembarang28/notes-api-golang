package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;uniqueIndex;not null"`
	Password  string    `gorm:"column:password;not null"`
	Photo     *string   `gorm:"column:photo"`
	createdAt time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
	Session   []Session `gorm:"foreignKey:user_id;references:id;OnDelete:CASCADE"`
	Notes     []Notes   `gorm:"foreignKey:user_id;references:id;OnDelete:CASCADE"`
}

func (u *User) TableName() string {
	return "users"
}

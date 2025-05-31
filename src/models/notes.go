package models

import "time"

type Notes struct {
	ID         string    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `gorm:"column:name;not null"`
	Notes      string    `gorm:"column:notes;not null;type:text"`
	CategoryID string    `gorm:"column:category_id;not null"`
	UserID     string    `gorm:"column:user_id;not null"`
	Tags       JSONB     `gorm:"column:tags;type:jsonb;not null;default:'[]'"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
	User       User      `gorm:"foreignKey:user_id;references:id;OnDelete:CASCADE"`
	Category   Category  `gorm:"foreignKey:category_id;references:id;OnDelete:CASCADE"`
}

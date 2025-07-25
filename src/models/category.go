package models

import "time"

type Category struct {
	ID          string    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string    `gorm:"column:name;not null;uniqueIndex"`
	UserID      string    `gorm:"column:user_id;not null"`
	Description string    `gorm:"column:description;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
	User        User      `gorm:"foreignKey:user_id;references:id;OnDelete:CASCADE"`
	Notes       []Notes   `gorm:"foreignKey:category_id;references:id;OnDelete:CASCADE"`
}

func (c *Category) TableName() string {
	return "category"
}

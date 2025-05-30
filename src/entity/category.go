package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          string  `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string  `gorm:"column:name;not null;uniqueIndex"`
	Description string  `gorm:"column:description;not null"`
	CreatedAt   string  `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt   string  `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
	Notes       []Notes `gorm:"foreignKey:category_id;references:id;OnDelete:CASCADE"`
}

func (c *Category) TableName() string {
	return "category"
}

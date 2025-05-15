package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int64          `gorm:"column:id;primaryKey"`
	Title       string         `gorm:"column:title;not null"`
	Description string         `gorm:"column:description"`
	CategoryID  uint           `gorm:"column:category_id;not null"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	UserID      string         `gorm:"column:user_id;not null"`
	Username    string         `gorm:"column:username;not null"`
	Quantity    int            `gorm:"column:quantity;not null"`
	Price       float64        `gorm:"column:price"`
	Image       string         `gorm:"column:image"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Product) TableName() string {
	return "products"
}

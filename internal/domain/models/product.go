package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int64          `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	CategoryID  uint           `json:"categoryId" gorm:"not null"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	UserID      string         `json:"userId" gorm:"not null"` 
	Username    string         `json:"userName" gorm:"not null"` 
	Quantity    int            `json:"quantity" gorm:"not null"`
	Price       float64        `json:"price,omitempty"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

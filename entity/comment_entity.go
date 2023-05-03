package entity

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	PhotoID   uint
	Message   string `gorm:"not null;type:text" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

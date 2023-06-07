package entity

import (
	"time"
)

type Socialmedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null;type:varchar(30)" validate:"required"`
	SocialMediaURL string `gorm:"not null;type:text" validate:"required"`
	UserID         uint   `gorm:"not null" validate:"required"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

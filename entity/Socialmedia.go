package entity

import (
	"time"
)

type Socialmedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null;type:varchar(30)" validate:"required"`
	SocialMediaURL string `gorm:"not null;type:text" validate:"required"`
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

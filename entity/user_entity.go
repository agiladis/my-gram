package entity

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null;unique;type:varchar(20)" validate:"required"`
	Email        string `gorm:"not null;unique;type:varchar(40)" validate:"required,email"`
	Password     string `gorm:"not null;type:varchar(80)" validate:"required"`
	Age          int    `gorm:"not null" validate:"required,gt=8"`
	Comments     []Comment
	Photos       []Photo
	SocialMedias []Socialmedia
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

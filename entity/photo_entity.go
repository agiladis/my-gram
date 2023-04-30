package entity

import "time"

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null;type:varchar(30)" validate:"required"`
	PhotoURL  string `gorm:"not null;type:text" validate:"required"`
	UserID    uint
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

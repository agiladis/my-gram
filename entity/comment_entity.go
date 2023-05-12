package entity

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null" validate:"required"`
	PhotoID   uint   `gorm:"not null" validate:"required"`
	Message   string `gorm:"not null;type:text" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentCreateRequest struct {
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
}

type CommentUpdateRequest struct {
	Message string `json:"message"`
}

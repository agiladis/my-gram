package entity

import "time"

type Photo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"not null;type:varchar(30)" validate:"required"`
	Caption   string `json:"caption" gorm:"type:text"`
	PhotoURL  string `json:"photo_url" gorm:"not null;type:text" validate:"required"`
	UserID    uint   `json:"user_id"`
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PhotoCreateRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type PhotoResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

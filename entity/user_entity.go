package entity

import "time"

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"not null;unique;type:varchar(20)" validate:"required"`
	Email        string `json:"email" gorm:"not null;unique;type:varchar(40)" validate:"required,email"`
	Password     string `json:"password" gorm:"not null;type:varchar(80)" validate:"required,min=6"`
	Age          int    `json:"age" gorm:"not null" validate:"required,gt=8"`
	Comments     []Comment
	Photos       []Photo
	SocialMedias []Socialmedia
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserAuthenticate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

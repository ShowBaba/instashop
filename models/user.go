package models

import (
	"time"
)

type User struct {
	ID             uint `gorm:"primaryKey"`
	Email          string
	FirstName      string
	LastName       string
	Password       string
	PhoneNumber    string
	ProfilePicture string
	IsVerified     *bool     `gorm:"default:false"`
	IsAdmin        *bool     `gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetUser struct {
	ID             uint `gorm:"primaryKey"`
	Email          string
	FirstName      string
	LastName       string
	PhoneNumber    string
	ProfilePicture string
	IsVerified     *bool     `gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (u *User) ToGetUser() *GetUser {
	return &GetUser{
		ID:             u.ID,
		Email:          u.Email,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		PhoneNumber:    u.PhoneNumber,
		ProfilePicture: u.ProfilePicture,
		IsVerified:     u.IsVerified,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

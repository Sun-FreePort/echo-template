package model

import (
	"time"
)

// User 用户
type User struct {
	Id              uint      `json:"id" gorm:"primaryKey;not null"`
	Name            string    `json:"name" gorm:"unique;not null"`
	Email           string    `gorm:"unique;not null"`
	EmailVerifiedAt time.Time `json:"email_verified_at" gorm:"default:true"`
	Password        string    `json:"password" gorm:"not null"`
	PhoneArea       int       `json:"phone_area"`
	PhoneNumber     string    `json:"phone_number"`
	RememberToken   string    `json:"remember_token" gorm:"default:true"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

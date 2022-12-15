package model

import "time"

// PersonalAccessToken Token
type PersonalAccessToken struct {
	Id            uint      `json:"id" gorm:"primaryKey;not null"`
	TokenableType string    `json:"tokenable_type"` // PHP 模型
	TokenableId   uint      `json:"tokenable_id"`   // 用户 ID
	Name          string    `json:"name"`           // 用户设备
	Token         string    `json:"token"`
	Abilities     string    `json:"abilities"`
	LastUsedAt    time.Time `json:"last_used_at" gorm:"default:true"`
	ExpiresAt     time.Time `json:"expires_at" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

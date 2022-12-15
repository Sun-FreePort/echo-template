package model

import (
	"time"
)

// TransGoods 跨城市货物运输
type TransGoods struct {
	Id        uint      `json:"id" gorm:"primaryKey;not null"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

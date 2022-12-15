package model

import (
	"time"
)

// Goods 持有物品
type Goods struct {
	Id        string    `json:"id"`
	UserId    uint      `json:"user_id"`
	Order     string    `json:"order"`   // 排序
	Index     uint      `json:"index"`   // dict Id
	Status    uint32    `json:"status"`  // 位状态，1 是否被装备
	Count     uint32    `json:"count"`   // 数量
	Wear      uint16    `json:"wear"`    // 耐久
	Effects   string    `json:"effects"` // 效果，JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

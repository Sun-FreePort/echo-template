package model

import (
	"time"
)

// Player 玩家本体
type Player struct {
	Id            uint      `json:"id" gorm:"primaryKey;not null"`
	Name          string    `json:"name"`
	Local         uint      `json:"local"`           // 地点 Id
	Hp            uint      `json:"hp"`              // 健康
	Thew          uint      `json:"thew"`            // 体力
	Enemy         uint      `json:"enemy"`           // 精力
	Attack        uint      `json:"attack"`          // 攻击力
	Defence       uint      `json:"defence"`         // 防御力
	Nimble        uint      `json:"nimble"`          // 敏捷
	HpMax         uint      `json:"hp_max"`          // 最大健康
	ThewMax       uint      `json:"thew_max"`        // 最大体力
	EnemyMax      uint      `json:"enemy_max"`       // 最大精力
	AttackMax     uint      `json:"attack_max"`      // 最大攻击力
	DefenceMax    uint      `json:"defence_max"`     // 最大防御力
	NimbleMax     uint      `json:"nimble_max"`      // 最大敏捷
	SlotHead      uint      `json:"slot_head"`       // 装备槽，里面放道具 Id
	SlotChest     uint      `json:"slot_chest"`      // 装备槽
	SlotHandLeft  uint      `json:"slot_hand_left"`  // 装备槽
	SlotHandRight uint      `json:"slot_hand_right"` // 装备槽
	SlotLeg       uint      `json:"slot_leg"`        // 装备槽
	SlotFootLeft  uint      `json:"slot_foot_left"`  // 装备槽
	SlotFootRight uint      `json:"slot_foot_right"` // 装备槽
	SlotAccessory uint      `json:"slot_accessory"`  // 装备槽
	RefreshedAt   uint      `json:"refreshed_at"`    // 刷新时间
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

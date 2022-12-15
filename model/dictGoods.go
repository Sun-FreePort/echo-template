package model

// DictGoods 字典：装备
type DictGoods struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	SkillTarget uint16 `json:"skill_target"` // 相关技能
	Wear        uint16 `json:"wear"`         // 耐久
	Effects     string `json:"effects"`      // 效果，JSON
	Type        uint16 `json:"type"`         // 类型，1 为通用品可叠加；2 为特殊品不可叠加；3 为左手装备不可叠加；4 为右手装备不可叠加；5 为双手装备不可叠加（置于左手，右手此时为 -1 Id）；6 为头部装备不可叠加；7 为胸腹装备不可叠加；8 为下身装备不可叠加；9 为左脚装备不可叠加；10 为右脚装备不可叠加；11 配饰装备不可叠加
}

package structure

import (
	"encoding/json"
	"errors"
	"github.com/Sun-FreePort/echo-template/cache"
	"github.com/Sun-FreePort/echo-template/model"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"strconv"
)

type DictGoods struct {
	Id      uint           `json:"id"`
	Name    uint           `json:"name"`
	Type    uint           `json:"type"`    // 类型，1 为通用品可叠加；2 为特殊品不可叠加；3 为左手装备不可叠加；4 为右手装备不可叠加；5 为双手装备不可叠加（置于左手，右手此时为 -1 Id）；6 为头部装备不可叠加；7 为胸腹装备不可叠加；8 为下身装备不可叠加；9 为左脚装备不可叠加；10 为右脚装备不可叠加；11 配饰装备不可叠加
	Wear    uint           `json:"wear"`    // 耐久
	Effects map[string]int `json:"effects"` // 效果，JSON
}

// GetByID 按 ID 获取本地配置
func (as *DictGoods) GetByID(id uint, redis *redis.Client, db *gorm.DB) (*DictGoods, error) {
	var result DictGoods
	str := cache.Get(redis, "dict:goods"+strconv.Itoa(int(id)))
	if str == "" {
		var m model.DictGoods
		err := db.Where(&model.DictGoods{Id: id}).First(&m).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}

			return nil, err
		}

		var effects map[string]int
		if err := json.Unmarshal([]byte(m.Effects), &effects); err != nil {
			return nil, err
		}
		result.Effects = effects
	} else {
		if err := json.Unmarshal([]byte(str), &result); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

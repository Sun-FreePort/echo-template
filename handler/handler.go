package handler

import (
	"github.com/Sun-FreePort/echo-template/cache"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Handler struct {
	db     *gorm.DB
	redis  *redis.Client
	userId uint
}

// NewHandler 注册全局变量
func NewHandler(redis *redis.Client, db *gorm.DB) *Handler {
	return &Handler{
		db:    db,
		redis: redis,
	}
}

// Register 注册路由
func (h *Handler) Register(router *echo.Group) {
	h.userId = 0
	h.AuthNot(router)

	auth := router.Group("")
	auth.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		tokenVal := cache.Get(h.redis, "token:"+key)
		if tokenVal != "" {
			i, err := strconv.Atoi(tokenVal)
			if err != nil {
				panic(err)
			}
			h.userId = uint(i)

			cache.Expire(h.redis, "token:"+key, time.Hour*48)

			// TODO 检查是否有工作，如果有，更新工作

			return true, nil
		}
		return false, nil
	}))
	h.AuthNeed(auth)
}

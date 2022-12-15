package handler

import "github.com/labstack/echo/v4"

func (h *Handler) AuthNot(router *echo.Group) {
	router.POST("/auth/signup", h.SignupUser)
	router.POST("/auth/login", h.LoginUser)
}

func (h *Handler) AuthNeed(router *echo.Group) {
	user := router.Group("/user")
	user.GET("/info", h.UserInfo)

	system := router.Group("/system")
	system.GET("/dict", h.SystemDict)
	system.GET("/dict/refresh", h.SystemDictRefresh)
}

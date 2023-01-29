package handler

import (
	"encoding/json"
	"github.com/Sun-FreePort/echo-template/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type systemDictRes struct {
	Goods  []model.DictGoods `json:"goods"`
	Errors map[string]string `json:"errors"` // 系统错误码和对应翻译
}

// SystemDict godoc
//
// @Summary     系统配置
// @Tags        system
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Success     200 {object} systemDictRes
// @Failure     403 {string} string "错误码"
// @Router      /system/dict [get]
func (h *Handler) SystemDict(c echo.Context) error {
	dict := new(systemDictRes)

	content, err := os.ReadFile("./config.json")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "dict error message not found")
	}

	var payload map[string]string
	err = json.Unmarshal(content, &payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "Error during JSON content")
	}
	dict.Errors = payload

	return c.JSON(http.StatusOK, dict)
}

func (h *Handler) SystemDictRefresh(c echo.Context) error {
	// TODO 清理整个 "dict:" Redis 缓存
	// TODO 将 JSON 文件读入变量
	// TODO 将变量按 “dict:{type}:{id}” 的方式存储起来

	return c.JSON(http.StatusOK, "{}")
}

package handler

import (
	"encoding/json"
	"github.com/Sun-FreePort/echo-game/cache"
	"github.com/Sun-FreePort/echo-game/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type newInfoRes struct {
	Token string `json:"token"`
}

func (h *Handler) NewInfo(c echo.Context) error {
	res := new(newInfoRes)
	return c.JSON(http.StatusCreated, res)
}

type userInfoRes struct {
	Ver    string        `json:"ver"`
	Change string        `json:"change"`
	Ts     int           `json:"ts"`
	Player model.Player  `json:"player"`
	Goods  []model.Goods `json:"goods"`
}

// UserInfo godoc
//
// @Summary     玩家信息
// @Tags        player
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Success     200 {object} userInfoRes
// @Failure     403 {string} string "错误码"
// @Router      /user/info [get]
func (h *Handler) UserInfo(c echo.Context) error {
	str := cache.Get(h.redis, "player:"+strconv.Itoa(int(h.userId)))
	res := userInfoRes{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return err
	}

	//slcB, _ := json.Marshal(playerInfo)

	return c.JSON(http.StatusCreated, res)
}

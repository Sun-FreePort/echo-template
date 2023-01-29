package handler

import (
	"errors"
	"github.com/Sun-FreePort/echo-template/cache"
	"github.com/Sun-FreePort/echo-template/help"
	"github.com/Sun-FreePort/echo-template/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type userAuthRes struct {
	Id    uint   `json:"id"`
	Token string `json:"token"`
}

type userLoginReq struct {
	Name     string `json:"name"`     // 昵称，仅限大小写英文、数字和下划线(_)
	Password string `json:"password"` // 密码，长度 8-32 位
	Device   string `json:"device"`   // 设备号，随意传，用于识别 Token 所属设备
}

// LoginUser godoc
//
// @Summary     用户登录
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       request body userLoginReq true "用户"
// @Success     200 {object} userAuthRes
// @Failure     403 {string} string "错误码"
// @Router      /auth/login [post]
func (h *Handler) LoginUser(c echo.Context) error {
	req := new(userLoginReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	res := new(userAuthRes)

	// 查询 DB 是否存在用户
	var user model.User
	err := h.db.First(&user, "name = ?", req.Name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusForbidden, "userNotHas")
	}
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// 校验用户密码
	if help.IHashCheckPassword(req.Password, user.Password) == false {
		return c.JSON(http.StatusUnauthorized, "authReject")
	}
	res.Id = user.Id
	res.Token = h.replaceToken(user.Id, req.Device)

	return c.JSON(http.StatusCreated, res)
}

type userSignupReq struct {
	Name        string `json:"name"`         // 昵称，仅限大小写英文、数字和下划线(_)
	Password    string `json:"password"`     // 密码，长度 8-32 位
	Email       string `json:"email"`        // 邮箱
	PhoneArea   int    `json:"phone_area"`   // 手机号所属国际编码，默认 86
	PhoneNumber string `json:"phone_number"` // 手机号码，最大 15 位
	Device      string `json:"device"`       // 设备号，随意传，用于识别 Token 所属设备
}

// SignupUser godoc
//
// @Summary     用户注册
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       request body userSignupReq true "用户"
// @Success     200 {object} userAuthRes
// @Failure     403 {string} string "错误码"
// @Router      /auth/signup [post]
func (h *Handler) SignupUser(c echo.Context) error {
	req := new(userSignupReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	res := new(userAuthRes)

	// 查询 DB 是否存在用户
	var user model.User
	err := h.db.First(&user, "name = ? OR email=?", req.Name, req.Email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
	} else if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusForbidden, "uniqueMust")
	}

	// 生成用户
	user = model.User{Name: req.Name, Password: help.IHashPassword(req.Password), Email: req.Email, PhoneArea: req.PhoneArea, PhoneNumber: req.PhoneNumber}
	err = h.db.Create(&user).Error
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	res.Id = user.Id
	res.Token = h.replaceToken(user.Id, req.Device)

	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) replaceToken(id uint, device string) string {
	key := help.IRandomString(28)

	// 存储 Token
	token := model.PersonalAccessToken{TokenableType: "App\\Models\\User", TokenableId: id, Name: device, Token: key, Abilities: "[\"*\"]"}
	err := h.db.Create(&token).Error
	if err != nil {
		return "Error1"
	}

	// 缓存新 Token 并将老 Token 注销
	var tokens []model.PersonalAccessToken
	err = h.db.Where("tokenable_id=?", id).Find(&tokens).Error
	if err != nil {
		return "Error2"
	}
	for _, v := range tokens {
		if v.Token != key {
			cache.Delete(h.redis, v.Token)
		}
	}
	cache.SetExpiration(h.redis, "token:"+key, strconv.Itoa(int(id)), time.Hour*48)

	return key
}

package server

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/component/config"
	"server/logic/orm/dal"
	"server/logic/password"
	"server/logic/token"
	"server/service/h"
	"time"
)

type LoginRequest struct {
	Password string `json:"password" binding:"required,len=40" label:"密码"`
	Phone    string `json:"phone" binding:"required,len=11" label:"手机号"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		h.Validator(c, err)
	} else {
		account, err := dal.Account.Where(dal.Account.Phone.Eq(loginRequest.Phone)).First()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.FailMessage(c, "账户不存在")
		} else if err != nil {
			h.Fail(c, err)
		} else {
			if len(account.Password) > 0 && password.Hash(loginRequest.Password, account.Salt) == account.Password {
				tokenString, err := token.Sign(&token.CustomClaims{
					Uid:      account.ID,
					Nickname: account.Nickname,
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
					},
				})
				if err != nil {
					h.Fail(c, err)
				} else {
					h.SetCookie(c, config.Config.GetString("server.clientTokenKey"), tokenString)
					h.OK(c, LoginResponse{
						Token: tokenString,
					})
				}
			} else {
				h.FailMessage(c, "密码错误")
			}
		}
	}
}

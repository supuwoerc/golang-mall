package server

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	helper_gen "github.com/langwan/langgo/helpers/gen"
	"gorm.io/gorm"
	"server/component/config"
	"server/logic/logic/password"
	"server/logic/logic/token"
	"server/logic/orm/dal"
	"server/logic/orm/model"
	"server/service/h"
	"time"
)

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Phone    string `json:"phone" binding:"required,len=11" label:"手机号"`
	Nickname string `json:"nickname" binding:"required,min=2,max=8" label:"昵称"`
	Password string `json:"password" binding:"required,len=40" label:"密码"` //sha1加密后传输
}

// RegisterResponse 注册响应信息
type RegisterResponse struct {
	Token string `json:"token"`
}

func Register(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		h.Validator(c, err)
	} else {
		fmt.Println(registerRequest)
		_, err := dal.Account.Where(dal.Account.Phone.Eq(registerRequest.Phone)).First()
		if err != gorm.ErrRecordNotFound {
			if err != nil {
				h.Fail(c, err)
			} else {
				h.ValidatorError(c, "phone", "电话号码已经被注册")
			}
		} else {
			salt, _ := helper_gen.RandString(16)
			hashPassword := password.Hash(registerRequest.Password, salt)
			account := model.Account{
				ID:       helper_gen.UuidShort(),
				Phone:    registerRequest.Phone,
				Nickname: registerRequest.Nickname,
				Password: hashPassword,
				Salt:     salt,
			}
			err := dal.Account.Create(&account)
			if err != nil {
				h.Fail(c, err)
			} else {
				var claims = token.CustomClaims{
					Uid:      account.ID,
					Nickname: account.Nickname,
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: time.Now().Add(time.Hour * 7 * 24).UnixMilli(),
					},
				}
				tokenString, err := token.Sign(&claims)
				if err != nil {
					h.Fail(c, err)
				} else {
					//将token种到客户端cookie里面，之后的客户端请求就能自动带上了
					h.SetCookie(c, config.Config.GetString("server.clientTokenKey"), tokenString)
					h.OK(c, tokenString)
				}
			}
		}
	}
}

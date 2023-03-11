package server

import (
	"github.com/gin-gonic/gin"
	"server/service/h"
)

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Phone    string `json:"phone" binding:"required,len=11" label:"手机号"`
	Nickname string `json:"nickname" binding:"required,min=2,max=8" label:"昵称"`
	Password string `json:"password" binding:"required,len=40" label:"密码"`
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
		h.OK(c, "hello register")
	}
}

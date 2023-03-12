package h

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/component/config"
	"server/logic/token"
)

// Auth 获取客户端在cookie存储的token做权限校验的中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(config.Config.GetString("server.clientTokenKey"))
		if err != nil {
			Forbidden(c) //取消
		} else {
			customClaims, err := token.Parse(tokenString)
			if err != nil {
				Forbidden(c)
			} else {
				c.Set("token", customClaims)
				//TODO:纪录访问到日志
				fmt.Println(customClaims.Nickname + "访问权限校验通过")
			}
		}
	}
}

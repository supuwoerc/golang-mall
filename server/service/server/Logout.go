package server

import (
	"github.com/gin-gonic/gin"
	"server/component/config"
	"server/service/h"
)

func Logout(c *gin.Context) {
	h.RemoveCookie(c, config.Config.GetString("server.clientTokenKey"))
	h.OK(c, nil)
}

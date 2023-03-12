package main

import (
	"github.com/gin-gonic/gin"
	_ "server/component/mysql" //初始化数据库
	"server/service/server"
)

func main() {
	g := gin.New()
	g.POST("/login", server.Login)
	g.POST("/register", server.Register)
	g.POST("/logout", server.Logout)
	_ = g.Run(":8080")
}

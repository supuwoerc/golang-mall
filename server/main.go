package main

import (
	"github.com/gin-gonic/gin"
	_ "server/component/mysql" //初始化数据库
)

func main() {
	g := gin.New()
	RouterRegister(g) //路由
	_ = g.Run(":8080")
}

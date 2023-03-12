package main

import (
	"github.com/gin-gonic/gin"
	_ "server/component/mysql" //初始化数据库
)

func main() {
	g := gin.New()
	RouterRegister(g)
	_ = g.Run(":8080")
}

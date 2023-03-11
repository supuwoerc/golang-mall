package main

import (
	"github.com/gin-gonic/gin"
	"server/service/server"
)

func main() {
	g := gin.New()
	//dsn := config.Config.GetString("mysql.dsn")
	//db, err := gorm.Open(mysql.Open(dsn))
	//if err != nil {
	//	panic(fmt.Errorf("创建数据库连接失败: %w", err))
	//}
	g.POST("/login", server.Login)
	g.POST("/register", server.Register)
	g.POST("/logout", server.Logout)
	//accountGroup := g.Group("/account")
	//accountGroup.POST("/profile", account.Profile)
	_ = g.Run(":8080")
}

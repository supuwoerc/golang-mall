package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/logic/orm/dal"
)

func main() {
	g := gin.New()
	const dsn = "golang_mall:password@tcp(localhost:3306)/golang_mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("创建数据库连接失败: %w", err))
	}
	dal.SetDefault(db)
	account, _ := dal.Account.Where(dal.Account.ID.Eq("1111")).First()
	fmt.Println(account)
	//g.POST("/login", server.Login)
	//g.POST("/register", server.Register)
	//g.POST("/logout", server.Logout)
	//accountGroup := g.Group("/account")
	//accountGroup.POST("/profile", account.Profile)
	g.Run(":8080")
}

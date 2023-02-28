package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/service/account/account"
	"server/service/server"
)

func main() {
	fmt.Println("hello")
	g := gin.New()
	g.POST("/login", server.Login)
	g.POST("/register", server.Register)
	g.POST("/logout", server.Logout)
	accountGroup := g.Group("/account")
	accountGroup.POST("/profile", account.Profile)
	g.Run(":8080")
}

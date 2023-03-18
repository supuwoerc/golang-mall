package main

import (
	"github.com/gin-gonic/gin"
	"server/service/account/goods"
	"server/service/h"
	"server/service/server"
)

func RouterRegister(g *gin.Engine) {
	v1 := g.Group("/api/v1")
	{
		v1.POST("/login", server.Login)
		v1.POST("/register", server.Register)
		v1.POST("/logout", server.Logout)
	}
	accountGroup := v1.Group("/account")
	accountGroup.Use(h.Auth()) //需要权限
	{
		goodGroup := accountGroup.Group("/goods")
		goodGroup.GET("/homepage", goods.HomePage)
		goodGroup.GET("/get", goods.Get)
	}
	orderGroup := accountGroup.Group("/order")
	//TODO:方法实现
	orderGroup.GET("/list", func(context *gin.Context) {

	})
}

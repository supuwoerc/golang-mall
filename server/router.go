package main

import (
	"github.com/gin-gonic/gin"
	"server/service/server"
)

func RouterRegister(g *gin.Engine) {
	v1 := g.Group("/api/v1")
	{
		v1.POST("/login", server.Login)
		v1.POST("/register", server.Register)
		v1.POST("/logout", server.Logout)
	}
}

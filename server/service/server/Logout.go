package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	c.String(http.StatusOK, "hello logout")
}

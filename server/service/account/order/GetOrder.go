package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrder(c *gin.Context) {
	c.String(http.StatusOK, "hello login")
}

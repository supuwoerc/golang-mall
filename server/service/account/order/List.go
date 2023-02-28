package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	c.String(http.StatusOK, "hello login")
}

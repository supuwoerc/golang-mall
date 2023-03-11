package h

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func OK(c *gin.Context, body any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"body":    body,
		"message": "",
	})
}
func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    0,
		"body":    nil,
		"message": err.Error(),
	})
}
func FailMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    0,
		"body":    nil,
		"message": msg,
	})
}
func FailCode(c *gin.Context, code int, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"body":    nil,
		"message": err.Error(),
	})
}
func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":    -1,
		"body":    nil,
		"message": "无权限访问",
	})
}

func Validator(c *gin.Context, err error) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		var result = gin.H{}
		for _, v := range errors {
			//validate.RegisterTagNameFunc将StructField改为取label
			result[v.StructField()] = v.Translate(Trans)
		}
		c.JSON(http.StatusUnprocessableEntity, result)
	} else {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
}

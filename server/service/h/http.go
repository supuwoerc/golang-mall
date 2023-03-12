package h

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"server/component/config"
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
		var errInfo = gin.H{}
		for _, v := range errors {
			//validate.RegisterTagNameFunc将StructField改为取label
			errInfo[v.StructField()] = v.Translate(Trans)
		}
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    -1,
			"body":    errInfo,
			"message": "表单验证失败",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    -1,
			"body":    err,
			"message": "表单验证失败",
		})
	}
}

func ValidatorError(c *gin.Context, tag, msg string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": -1,
		"body": gin.H{
			tag: msg,
		},
		"message": "字段校验失败",
	})
}

func SetCookie(c *gin.Context, key, val string) {
	c.SetCookie(key, val, 365*86400, "/", config.Config.GetString("server.host"), false, false)
}

func RemoveCookie(c *gin.Context, key string) {
	c.SetCookie(key, "", -1, "/", config.Config.GetString("server.host"), false, true)
}

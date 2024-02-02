package utils

import (
	"crypto/sha256"
	"fmt"

	"github.com/gin-gonic/gin"
)

var DateTimeFormat = "2006-01-02 15:04:05"

// sha256 加密, 用于加密密码
func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

// JSON 返回
func Res(c *gin.Context, code int, msg string, data any) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(200, gin.H{
		"status": 1,
		"code":   code,
		"msg":    msg,
		"data":   data,
	})
}

func Error(c *gin.Context, msg string) {
	Res(c, 500, msg, nil)

}

package controller

import (
	"gin-sites/service"
	"gin-sites/utils"

	"github.com/gin-gonic/gin"
)

// 检验 token 是否有效
// 经过 jwt middleware 后到达的请求中 token 是有效的,此函数主要验证数据库中用户是否存在
func CheckTokenHandle(c *gin.Context) {
	username, nameExist := c.Get("username")
	userID, idExist := c.Get("userID")

	if !nameExist || !idExist {
		utils.Res(c, 401, "用户信息不存在", nil)
		return
	}

	tmp := service.IsValidUser(username.(string), userID.(int32))
	if tmp.Status == 0 {
		c.JSON(resBad(tmp.Msg))
		return
	}
	c.JSON(resOK())
}

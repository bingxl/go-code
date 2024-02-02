package controller

import (
	"gin-sites/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户登录
func LoginHandle(c *gin.Context) {

	// 定义接收的参数, 并做validate
	// binding 写法 https://github.com/go-playground/validator
	// 自定义validate https://gin-gonic.com/zh-cn/docs/examples/custom-validators/
	type userLogin struct {
		Username string `form:"username" json:"username" binding:"required"`
		Pwd      string `form:"password" json:"password" binding:"required"`
	}
	var user userLogin

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(resBad("请求参数错误" + err.Error()))
		return
	}

	c.JSON(res(service.Authorization(user.Username, user.Pwd)))

}

// 获取用户详情
func GetUserHandle(c *gin.Context) {

	// jwt middleware 中设置的当前请求token中的 用户名和id
	curUserId, uidExists := c.Get("userID")

	// 路由参数中的用户id
	userId, err := strconv.Atoi(c.Param("userId"))

	// 参数不存在,或者 userId 转换失败 或者 userId == 0(0值)
	if !uidExists || err != nil || userId == 0 {
		c.JSON(resBad("用户名或id不存在"))
		return
	}
	tmpRet := service.GetUserDetail(curUserId.(int32), int32(userId))
	c.JSON(res(tmpRet))

}

// 创建用户
func CreateUserHandle(c *gin.Context) {
	type userT struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Age      int32  `json:"age"`
	}

	user := userT{}
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(resBad(err.Error()))
		return
	}

	re := service.CreateUser(user.Username, user.Password, user.Email, user.Age)
	c.JSON(res(re))
}

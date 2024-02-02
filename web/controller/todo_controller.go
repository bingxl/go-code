package controller

import (
	"strconv"

	"gin-sites/service"

	"github.com/gin-gonic/gin"
)

// 获取待办事项列表
func TodoListHandle(c *gin.Context) {
	curUserId, exists := c.Get("userID")
	if !exists {
		c.JSON(resBad("userId error"))
		return
	}
	// page 和 limit 设置
	// http:/xxx/sxx?page=0&limit=20
	page, pageErr := strconv.Atoi(c.DefaultQuery("page", "0"))
	limit, limitErr := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if pageErr != nil {
		page = 0
	}
	if limitErr != nil {
		limit = 20
	}

	tmp := service.GetUserAllTodo(curUserId.(int32), page, limit)
	c.JSON(res(tmp))
}

// 添加一个待办事项列表
func TodoAddHandle(c *gin.Context) {
	type todoParamsT struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Tag     string `json:"tag"`
		Owner   int32  `json:"owner"`
	}
	todo := todoParamsT{}
	getUID, exists := c.Get("userID")
	if !exists {
		c.JSON(resBad("userID error"))
		return
	}
	todo.Owner = getUID.(int32)

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(resBad(err.Error()))
		return
	}

	tmp := service.CreateTodo(todo.Title, todo.Content, todo.Tag, todo.Owner)
	c.JSON(res(tmp))
}

// 删除一个待办事项
func TodoDeleteHandle(c *gin.Context) {
	todoId := c.Param("todoId")
	curUserId := c.GetInt("userID")
	tmp := service.DeleteTodo(todoId, int32(curUserId))
	c.JSON(res(tmp))

}

// 更新一个待办事项
func TodoUpdateHandle(c *gin.Context) {
	type todoParamsT struct {
		ID      string `json:"id" binding:"required"`
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Status  int    `json:"status"`
		Tag     string `json:"tag"`
	}
	todo := todoParamsT{}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(resBad(err.Error()))
		return
	}

	// 参数验证
	if todo.ID == "" || todo.Title == "" || todo.Content == "" {
		c.JSON(resBad("id, title, content 都不能为空"))
		return
	}

	tmp := service.UpdateTodo(todo.Title, todo.Content, todo.Tag, int32(todo.Status))
	c.JSON(res(tmp))

}

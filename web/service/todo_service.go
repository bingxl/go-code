package service

import (
	"gin-sites/dao"
	"gin-sites/dao/model"
)

var todoDao = dao.NewTodoDao()

// 创建一个todo
// @param status int32 任务状态 0 | 1, 0 未完成, 1 已完成
func CreateTodo(title, content, tag string, owner int32) ResT {
	todoId, err := todoDao.Create(&model.Todo{
		Title:   title,
		Content: content,
		Tag:     tag,
		Owner:   owner,
		Status:  0,
	})

	if err != nil {
		return fail(err.Error())
	}

	return success(todoId)
}

// 更新todo
func UpdateTodo(title, content, tag string, status int32) ResT {
	err := todoDao.UpdateTodo(&model.Todo{
		Title:   title,
		Content: content,
		Tag:     tag,
		Status:  status,
	})
	if err != nil {
		return fail(err.Error())
	}
	return success(nil)
}

// 获取用户的待办列表
func GetUserAllTodo(owner int32, page, limit int) ResT {
	all, err := todoDao.GetAllTodoByOwner(owner, page, limit)
	if err != nil {
		return fail(err.Error())
	}
	return success(all)
}

// 删除待办事项
func DeleteTodo(todoId string, curUserId int32) ResT {
	// @TODO 校验权限
	if err := todoDao.DeleteById(todoId); err != nil {
		return fail(err.Error())
	}
	return success(nil)

}

package dao

import (
	"errors"
	"gin-sites/dao/model"

	"github.com/google/uuid"
)

type todoDao struct {
}

func NewTodoDao() *todoDao {
	return &todoDao{}
}

// 创建一个 TODO 返回创建后的TODO ID 和错误
func (t *todoDao) Create(todo *model.Todo) (ID string, err error) {
	var _todo = dao.Todo
	var todoCtx = _todo.WithContext(ctx)

	id, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Join(err, errors.New("生成uuid失败"))
	}
	todo.ID = id.String()
	err = todoCtx.Create(todo)

	return todo.ID, err
}

// 获取用户的所有待办列表
func (t *todoDao) GetAllTodoByOwner(uid int32, page int, limit int) ([]*model.Todo, error) {
	var _todo = dao.Todo
	var todoCtx = _todo.WithContext(ctx)
	return todoCtx.Where(_todo.Owner.Eq(uid)).Offset((page - 1) * limit).Limit(limit).Find()
}

// 更新 todo 但todo.ID 和todo.Owner 不会更新
func (t *todoDao) UpdateTodo(todo *model.Todo) error {
	var _todo = dao.Todo
	var todoCtx = _todo.WithContext(ctx)
	_, err := todoCtx.Omit(_todo.ID, _todo.Owner).Where(_todo.ID.Eq(todo.ID)).Updates(todo)
	return err
}

// 软删除
func (t *todoDao) DeleteById(todoId string) error {
	var _todo = dao.Todo
	var todoCtx = _todo.WithContext(ctx)

	_, err := todoCtx.Where(_todo.ID.Eq(todoId)).Delete()
	return err
}

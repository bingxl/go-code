package dao

import (
	"errors"
	"fmt"
	"gin-sites/dao/model"

	"gorm.io/gorm"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}

// 创建用户
func (uDo *UserDao) Create(username, pwd, email string, age int32) (int32, error) {
	// dao 定义在 init.go 文件
	var _user = dao.User
	var userCtx = _user.WithContext(ctx)

	// 先从数据库找找
	_, err := uDo.FindByName(username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrUsernameExists
	}

	tmpU := &model.User{Username: username, Password: pwd, Email: email, Age: age}
	err = userCtx.Select(_user.Username, _user.Password, _user.Email, _user.Age).Create(tmpU)
	return tmpU.ID, err
}

// 查找用户名 为 username 的第一个用户
func (uDo *UserDao) FindByName(username string) (*model.User, error) {
	// dao 定义在 init.go 文件
	var _user = dao.User
	var userCtx = _user.WithContext(ctx)

	return userCtx.Where(_user.Username.Eq(username)).First()
}

func (uDo *UserDao) FindByNameAndID(username string, id int32) (*model.User, error) {
	u := dao.User
	return u.WithContext(ctx).Where(u.Username.Eq(username), u.ID.Eq(id)).First()
}

// 依据用户名和密码查找
func (uDo *UserDao) FindByNameAndPwd(username, pwd string) (*model.User, error) {
	// dao 定义在 init.go 文件
	var _user = dao.User
	var userCtx = _user.WithContext(ctx)
	fmt.Println("in /dao/ dao is: ", dao)

	return userCtx.Where(_user.Username.Eq(username), _user.Password.Eq(pwd)).First()
}
func (uDo *UserDao) FindById(uid int32) (*model.User, error) {
	// dao 定义在 init.go 文件
	var _user = dao.User
	var userCtx = _user.WithContext(ctx)

	return userCtx.Where(_user.ID.Eq(uid)).First()
}

// 更新用户信息, uid 不能更改, username 会查重
func (uDo *UserDao) Update(uid int32, username, pwd, email string, age int32) error {
	// dao 定义在 init.go 文件
	var _user = dao.User
	var userCtx = _user.WithContext(ctx)

	// 查询数据库 中是否有其他用户名为username
	_, err := userCtx.Where(_user.ID.Neq(uid), _user.Username.Eq(username)).First()

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUsernameExists
	}

	// 使用结构体更新时会忽略 field 的 0 值, 但会更新 updated_at
	_, err = userCtx.Where(_user.ID.Eq(uid)).Updates(&model.User{Username: username, Password: pwd, Email: email, Age: age})
	return err
}

// 获取所有有效(未被软删除)用户信息
func (uDo *UserDao) GetAllUsers() ([]*model.User, error) {
	// dao 定义在 init.go 文件
	var _user = dao.User
	var userCtx = _user.WithContext(ctx)

	return userCtx.Find()
}

// 软删除, 更新 deleted_at 字段(如果有的话), 没有此字段时 是直接从数据库中删除
func (uDo *UserDao) Delete(uid int32) error {
	// dao 定义在 init.go 文件
	var _user = dao.User
	var userCtx = _user.WithContext(ctx)

	_, err := userCtx.Where(_user.ID.Eq(uid)).Delete()
	return err
}

// 从数据库中删除用户
func (uDo *UserDao) DeleteFromDb(uid int32) error {
	// dao 定义在 init.go 文件
	var _user = dao.User
	var userCtx = _user.WithContext(ctx)

	_, err := userCtx.Where(_user.ID.Eq(uid)).Unscoped().Delete()
	return err
}

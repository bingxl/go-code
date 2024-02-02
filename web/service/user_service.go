package service

import (
	"errors"
	"gin-sites/dao"
	"gin-sites/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var userDao = dao.NewUserDao()

// 鉴权处理, (获取token), 参数在 handle 里校验, 须保证参数有效
func Authorization(username, pwd string) ResT {
	user, err := userDao.FindByNameAndPwd(username, utils.Sha256(pwd))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ResT{
			http.StatusNotFound,
			0,
			"用户名或密码错误",
			nil,
		}
	}

	token := utils.Generate(user.Username, user.ID, false)

	return ResT{
		http.StatusOK,
		1,
		"获取成功",
		gin.H{
			"token": token,
		},
	}
}

// 获取targetUid 用户的详细信息 omit: password, deleted_at
func GetUserDetail(curUid int32, targetUid int32) ResT {

	// @TODO 权限校验, 校验当前用户是否有权获取 targetUid 的信息

	u, err := userDao.FindById(targetUid)
	if err != nil {
		return ResT{
			Code:   500,
			Status: 0,
			Msg:    err.Error(),
			Data:   nil,
		}
	}

	return ResT{
		Code:   200,
		Status: 1,
		Msg:    "success",
		Data: map[string]any{
			"id":         u.ID,
			"age":        u.Age,
			"username":   u.Username,
			"email":      u.Email,
			"created_at": u.CreatedAt,
			"updated_at": u.UpdatedAt,
		},
	}
}

// 创建用户
func CreateUser(username, pwd, email string, age int32) ResT {
	uid, err := userDao.Create(username, pwd, email, age)
	if err != nil {
		return ResT{
			Code:   400,
			Status: 0,
			Msg:    "创建用户失败" + err.Error(),
		}
	}

	return ResT{
		Code:   200,
		Status: 1,
		Msg:    "success",
		Data:   uid,
	}
}

func IsValidUser(username string, id int32) ResT {
	u, err := userDao.FindByNameAndID(username, id)
	return handle(u, err)
}

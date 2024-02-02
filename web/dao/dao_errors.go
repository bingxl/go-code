package dao

import "errors"

// dao 包中用到的 自定义 error, 除此外还会用到 gorm.io/gorm 包中的error
var (
	ErrUsernameExists = errors.New("用户名已存在") // 创建用户时用户名已存在
)

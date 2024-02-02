// 初始化数据库连接
package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gin-sites/dao/query"
)

// DB 仅在 dao 包中使用, 与数据库的所有操作都在 dao  包中
var DB *gorm.DB
var err error
var dao *query.Query
var ctx = context.Background()

func init() {

	user := "root"
	password := "lxb106170$$"
	addr := "127.0.0.1"
	port := 3306
	dbName := "gin"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, addr, port, dbName)

	DB, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{

			// 表明命名决策
			NamingStrategy: schema.NamingStrategy{
				// 单表名
				SingularTable: true,
			},

			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			}),
		},
	)
	dao = query.Use(DB)

	if err != nil {
		log.Fatalln("数据库连接错误", err)

	}

}

// 关闭 DB 连接
func Close() {
	if DB != nil {
		sqlDb, err := DB.DB()
		if err != nil {
			sqlDb.Close()
		}
	}
}

package model

import (
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {

	db := newConn(connString)
	//禁止表名复数形式
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// 自定义的日志打印sql语句
	db.SetLogger(&GormLogger{})

	db.DB().SetMaxOpenConns(100)                 //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(20)                  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	db.DB().SetConnMaxLifetime(time.Second * 30) //超时

	DB = db

	migration()
}

//newConn 创建新连接
func newConn(connString string) *gorm.DB {
	// logger.Info(connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	return db
}

//GetDB 获取gorm.DB对象
func GetDB() *gorm.DB {
	return DB
}

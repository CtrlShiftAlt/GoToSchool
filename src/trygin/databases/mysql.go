package databases

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Db .
var Db *gorm.DB

var (
	host     = "127.0.0.1"
	post     = "3306"
	username = "root"
	password = "root"
	database = "test"
	charset  = "utf8mb4"
)

func init() {
	// 实例 Mysql
	var err error
	dsn := username + ":" + password + "@tcp(" + host + ":" + post + ")/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		
	})
	if err != nil {
		log.Panicln("err: ", err.Error())
	}
	sqlDB, err := Db.DB()
	if err != nil {
		log.Panicln("err: ", err.Error())
	}
	// 空闲连接池中连接最大数量
	sqlDB.SetMaxIdleConns(10)
	// 打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}

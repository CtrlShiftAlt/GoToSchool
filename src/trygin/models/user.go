package models

import (
	"fmt"
	"time"
)

// User 会员表
type User struct {
	UID       uint   `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar(20);not null;comment:用户名"`
	Password  string `gorm:"type:char(32);not null;comment:密码"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (U *User) hasTable() {
	// 如果有缓存，则先记录表名至缓存，缓存中存在表名则不需以下操作
	if !db.Migrator().HasTable(U) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})
		// db.Table(userAPILog).Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&UserAPILog{})
	}
}

// AddUser 添加用户
func (U *User) AddUser(username string, password string) {
	U.hasTable()
	userData := User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&userData)
}

// FindUser 添加用户
func (U *User) FindUser(username string) User {
	U.hasTable()
	data := User{}
	result := db.Where(&User{Username: username}).First(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return data
}

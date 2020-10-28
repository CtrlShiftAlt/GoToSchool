package models

import (
	"time"
)

// Book .
type Book struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(10);not null;comment:名字"`
	Name2     string `gorm:"not null"`
	Name3     string `gorm:"default:;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// tableName 表名
var book string

func (B *Book) hasTable() {
	book = "book" + time.Now().Format("_20060102")
	if !db.Migrator().HasTable(book) {
		db.Table(book).Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&Book{})
		// db.Table(tableName).Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Book{})
	}
}

//AddBook .
func (B *Book) AddBook() {
	B.hasTable()
	bookData := Book{
		Name:      "name",
		Name2:     "name2",
		Name3:     "name3",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Table(book).Create(&bookData)
}

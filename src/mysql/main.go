package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	id   int
	name string
	age  int
}

type student struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var db *sql.DB // 是一个连接池对象
var xdb *sqlx.DB

// Go链接MySQL
func initDB() (err error) {
	// 数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	db, err = sql.Open("mysql", dsn) // 不会校验用户名密码是否正确
	if err != nil {                  // dsn 格式不正确时报错
		return
	}
	err = db.Ping() // 尝试连接 验证账号密码
	if err != nil {
		return
	}
	return
}

// queryRow 查一行
func queryRow() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.QueryRow(sqlStr, 2).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("Scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
}

// queryMultiRow 查多行
func queryMultiRow() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func insert() {
	sqlStr := "insert into user(name, age) values (?, ?)"
	ret, err := db.Exec(sqlStr, "bibu", 28)
	if err != nil {
		fmt.Printf("insert faleld, err: %v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func update() {
	sqlStr := "update user set name = ? where id = ?"
	ret, err := db.Exec(sqlStr, "good", 1)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 影响行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)

}

// prepareQuery 预处理
func prepareQuery() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err: %v\n", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	defer rows.Close()

	// 循环读取结果中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func prepareInsert() {
	sqlStr := "insert into user(name, age) values (?, ?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err: %v\n", err)
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec("july", 25)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func prepareUpdate() {
	sqlStr := "update user set name = ? where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare failed, err: %v\n", err)
		return
	}
	defer stmt.Close()

	ret, err := stmt.Exec("Sala", 4)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 影响行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

func prepareDeleteRow() {
	sqlStr := "delete from user where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare failed, err: %v\n", err)
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec(1)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}

// transaction 事务
func transaction() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age = 29 where id = ?"
	_, err = tx.Exec(sqlStr1, 4)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	sqlStr2 := "Update user set age = 28 where id = ?"
	_, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed, err: %v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}

// sqlx
func iniDBsqlx() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	xdb, err = sqlx.Connect("mysql", dsn) // 注意作用域
	if err != nil {
		fmt.Printf("connect DB failed, err: %v\n", err)
		return
	}

	err = xdb.Ping()
	if err != nil {
		fmt.Printf("ping sqlx DB failed, err: %v\n", err)
		return
	}
	xdb.SetMaxOpenConns(20)
	xdb.SetMaxIdleConns(10)
	return
}

func queryRowsqlx() {
	var us student
	err := xdb.Get(&us, "select id, name, age from user where id = ?", 2)
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", us)
}

func queryMultiRowsqlx() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []student
	err := xdb.Select(&users, sqlStr, 0)
	if err == sql.ErrNoRows {
		fmt.Printf("not found data of the id:%d", 1)
	}
	if err != nil {
		fmt.Printf("query failed, err %v\n", err)
		return
	}
	fmt.Printf("users: %#v\n", users)
}

func insertsqlx() {
	sqlStr := "insert into user(name, age) values (?, ?)"
	ret, err := xdb.Exec(sqlStr, "slqx", 5)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", theID)
}

func uodatesqlx() {
	sqlStr := "update user set name = ? where id = ?"
	ret, err := xdb.Exec(sqlStr, "Tank", 6)
	if err != nil {
		fmt.Printf("Exec failed,err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success, afected rows: %d\n", n)
}

func deletesqlx() {
	sqlStr := "delete from user where id = ?"
	ret, err := xdb.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed,err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("delete failed,err: %v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}

func transactionsqlx() {
	tx, err := xdb.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed,err %v\n", err)
		return
	}
	sqlStr1 := "update user set age = 8 where id = ?"
	tx.Exec(sqlStr1, 2)
	sqlStr2 := "update user set age = 9 where id = ?"
	tx.Exec(sqlStr2, 4)
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed, err: %v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}

func main() {
	// err := initDB()
	// if err != nil {
	// 	fmt.Println("init DB failed, err: ", err)
	// }
	// fmt.Println("MySQL 连接成功！")

	// queryRow()
	// queryMultiRow()
	// insert()
	// update()
	// deleteRowDemo()

	// prepareQuery()
	// prepareInsert()
	// prepareUpdate()
	// prepareDeleteRow()

	// transaction()

	err := iniDBsqlx()
	if err != nil {
		fmt.Println("init sqlx DB failed, err: ", err)
	}
	fmt.Println("sqlx 连接成功！")
	defer xdb.Close()

	// queryRowsqlx()
	// queryMultiRowsqlx()
	// insertsqlx()
	// uodatesqlx()
	// deletesqlx()

	// transactionsqlx()
}

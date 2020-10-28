package models

import (
	"database/sql"
	"fmt"
	"time"
)

// UserAPILog .
type UserAPILog struct {
	ID         uint32 `gorm:"primaryKey"`
	UID        uint32 `gorm:"not null"`
	Name       string `gorm:"type:varchar(20);not null"`
	IP         string `gorm:"type:varchar(20);not null;comment:访问ip"`
	URI        string `gorm:"type:varchar(64);not null;comment:api地址"`
	Params     string `gorm:"type:varchar(32);not null;default:'';comment:参数"`
	LogContent string `gorm:"type:text;comment:日志内容"`
	ActivedAt  sql.NullTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// BeforeCreate 钩子
// func (U *UserAPILog) BeforeCreate(*gorm.DB) (err error) {
// 	U.hasTable()
// 	return nil
// }

// userAPILog 表名 // 默认查找当天的记录，查找多日的需根据表名查找
var userAPILog string

func (U *UserAPILog) hasTable() {
	userAPILog = "user_api_log" + time.Now().Format("_200601")
	// 如果有缓存，则先记录表名至缓存，缓存中存在表名则不需以下操作
	if !db.Migrator().HasTable(userAPILog) {
		db.Table(userAPILog).Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&UserAPILog{})
		// db.Table(userAPILog).Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&UserAPILog{})
	}
}

// Find .
func (U *UserAPILog) Find() []UserAPILog {
	U.hasTable()
	var data []UserAPILog
	result := db.Table(userAPILog).Find(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return data
}

// FindAll .
func (U *UserAPILog) FindAll() {
	U.hasTable()
	rows, err := db.Table(userAPILog).Rows()
	if err != nil {
		fmt.Println(err)
	}
	var id, uid, name, ip, uri, params, logContent, activedAt, createdAt, updatedAt string
	for rows.Next() {
		err = rows.Scan(&id, &uid, &name, &ip, &uri, &params, &logContent, &activedAt, &createdAt, &updatedAt)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, uid, name, ip, uri, params, logContent, activedAt, createdAt, updatedAt)
	}
}

// AddLog 写日志
func (U *UserAPILog) AddLog(uid uint32, content string) {
	U.hasTable()
	logData := UserAPILog{
		UID:        uid,
		Name:       "Ruach",
		IP:         "192.168.101.99",
		URI:        "192.168.101.99:8084/h5/friend/getlist",
		Params:     "{limit:1}",
		LogContent: content,
		ActivedAt:  sql.NullTime{Time: time.Now(), Valid: true},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	result := db.Table(userAPILog).Create(&logData)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

// func test() {
// 	//同步字段类型、注释，默认关闭
// 	_syncComment := false
// 	if _syncComment {
// 		_regComment := regexp.MustCompile(`comment\('([^']+)'\)`)
// 		// _regType := regexp.MustCompile(`[\s]*((int|smallint|varchar)\([\d]+\))\s+`)
// 		// _regType2 := regexp.MustCompile(`[\s]*(text)\s+`)
// 		// _session := engine.NewSession()

// 		// defer _session.Close()
// 		_slice := [0]string{}
// 		for _, _v := range _slice {
// 			_t := reflect.TypeOf(_v).Elem()
// 			fieldNum := _t.NumField()
// 			k := 0
// 			for i := 0; i < fieldNum; i++ {
// 				if k > 1000 {
// 					break
// 				}
// 				k = k + 1
// 				//
// 				_f := _t.Field(i)
// 				_xorm := _f.Tag.Get("xorm")
// 				if _xorm == "extends" {

// 				}
// 				_matchesComment := _regComment.FindAllStringSubmatch(_xorm, -1)
// 				if len(_matchesComment) > 0 {
// 					// _typeStr := ""
// 					// _matchesType := _regType.FindAllStringSubmatch(_xorm, -1)
// 					// if len(_matchesType) > 0 {
// 					// 	_typeStr = _matchesType[0][1]
// 					// } else {
// 					// 	_matchesType2 := _regType2.FindAllStringSubmatch(_xorm, -1)
// 					// 	if len(_matchesType2) > 0 {
// 					// 		_typeStr = _matchesType2[0][1]
// 					// 	}
// 					// }
// 					//
// 					// _sqlSlice := []string{
// 					// 	"alter table",
// 					// 	"`ppospro_" + utils.SnakeString(_t.Name()) + "`",
// 					// 	"modify column",
// 					// 	"`" + utils.SnakeString(_f.Name) + "`",
// 					// 	_typeStr,
// 					// 	"comment",
// 					// 	"'" + _matchesComment[0][1] + "'",
// 					// }
// 					// _sql := strings.Join(_sqlSlice, " ")
// 					// _, _err := _session.Exec(_sql)
// 					// fmt.Println(_sql, _err)
// 				}
// 			}
// 		}
// 	}
// }

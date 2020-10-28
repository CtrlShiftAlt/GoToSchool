package iniconf

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// MySQLConfig .
type MySQLConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig .
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

// Config .
type Config struct {
	MySQLConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 参数的校验
	// 传进来的data参数必须是指针 需要在函数中赋值
	t := reflect.TypeOf(data)
	// fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		// err = errors.New("data should be a pointer")
		err = fmt.Errorf("data param should be a pointer")
		return
	}
	// 传进来的data参数必t:=reflect.TypeOf(data)须是结构体类型指针 配置文件的各种键值对需要赋值给结构体字段
	if t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("data param should be a struct")
		return
	}
	// 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// string(b) // 字节类型文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\r\n")
	// fmt.Println(lineSlice)
	// 一行一行的读数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首位空格
		line = strings.TrimSpace(line)
		// 如果是注释或空行就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		// 如果\[开头的就表示是节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d systax error", idx+1)
				return
			}
			// 把这一行首尾的\[\]去掉，读取到中奖的内容把首位的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d systax error", idx+1)
				return
			}

			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找了对应的嵌套结构体,把字段名记下来
					structName = field.Name
					fmt.Printf("扎到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 如果不是\[开头的就是=分割的键值对
			// 以等号分割一行，等号左边是key 等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d systax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			// 根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			// 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // tag信息是存储在类型信息中
				fileType = field
				if field.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = field.Name
					break
					// sValue.Elem().SetString(value)
				}
			}
			// 如果key = tag 给这个字段赋值
			// 根据fieldName去取出这个字段
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字符
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			// 对其赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}
	}
	return
}

func main() {
	var cfg Config
	// var x = new(int)
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}

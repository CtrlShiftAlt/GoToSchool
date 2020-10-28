package iniconf

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
)

// IniTitle .
type IniTitle struct {
	MYSQL IniMySQL `tag:"IniMySQL"`
	REDIS IniRedis `tag:"IniRedis"`
}

var iniTitleObj = IniTitle{}
var typeObj reflect.Type
var valueObj reflect.Value
var ptrValueObj reflect.Value

// setReflectTypeObj ..
func setReflectTypeObj(tag string) bool {
	switch tag {
	case "IniMySQL":
		iniTitleObj.MYSQL = IniMySQL{}
		typeObj = reflect.TypeOf(iniTitleObj.MYSQL)
		valueObj = reflect.ValueOf(iniTitleObj.MYSQL)
		ptrValueObj = reflect.ValueOf(&iniTitleObj.MYSQL)
		return true
	case "IniRedis":
		iniTitleObj.REDIS = IniRedis{}
		typeObj = reflect.TypeOf(iniTitleObj.REDIS)
		valueObj = reflect.ValueOf(iniTitleObj.REDIS)
		ptrValueObj = reflect.ValueOf(&iniTitleObj.REDIS)
		return true
	}
	return false
}

// LoadFile ...
func LoadFile(fp string, tt string) interface{} {
	tt = strings.ToUpper(tt) // MYSQL

	// file data
	var data []byte
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Printf("File reading error, err: %v\n", err)
		return nil
	}
	// file data array
	file := strings.Split(string(data), "\r\n")
	// iniTitleObj reflect object
	refTypeTitleObj := reflect.TypeOf(iniTitleObj)

	for _, line := range file {
		// [TITLE]
		titleRegl := regexp.MustCompile(`^\[(\w+)\]\s*$`)
		if titleRegl == nil {
			fmt.Printf("title regexp err\n")
			return nil
		}
		titleStringArr := titleRegl.FindStringSubmatch(line)
		if len(titleStringArr) == 2 {
			t1, ok := refTypeTitleObj.FieldByName(titleStringArr[1])
			if !ok {
				fmt.Printf("FieldByName %s err\n", titleStringArr[1])
				return nil
			}

			tag, _ := t1.Tag.Lookup("tag") // IniMySQL IniRedis

			if ok = setReflectTypeObj(tag); !ok {
				fmt.Printf("Unfind type %s\n", tag)
				return nil
			}
			continue
		}

		itemsRegl := regexp.MustCompile(`^(\w+)\s*=\s*([\w|"|.]+)\s*$`)
		if itemsRegl == nil {
			fmt.Printf("item regexp err\n")
			return nil
		}
		itemsStringArr := itemsRegl.FindStringSubmatch(line)
		if len(itemsStringArr) == 3 {
			num := typeObj.NumField()
			for i := 0; i < num; i++ {
				if typeObj.Field(i).Tag.Get("tag") == itemsStringArr[1] {
					ptrValueObj.Elem().Field(i).SetString(itemsStringArr[2])
				}
			}
			continue
		}
		// 
	}

	switch tt {
	case "MYSQL":
		return iniTitleObj.MYSQL
	case "REDIS":
		return iniTitleObj.REDIS
	}
	return nil
}

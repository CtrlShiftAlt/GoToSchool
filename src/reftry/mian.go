package main

import (
	"fmt"
	"reflect"
)

// Student .
type Student struct {
	ID   int    `json:"id" tag:"id"`
	Name string `json:"name" tag:"id"`
}

// Hello .
func (s Student) Hello() {
	fmt.Println("我是一个学生")
}

// People .
type People struct {
	Student
}

func main() {
	// var name string = "咖啡色的羊驼"
	// reflectType := reflect.TypeOf(name)
	// reflectValue := reflect.ValueOf(name)
	// fmt.Println(reflectType)
	// fmt.Println(reflectValue)

	// s := Student{ID: 1, Name: "咖啡色的羊驼"}
	// t := reflect.TypeOf(s)
	// fmt.Println(t.Name(), t.Kind())

	// fmt.Println(t.NumMethod())
	// for i := 0; i < t.NumMethod(); i++ {
	// 	m := t.Method(i)
	// 	fmt.Printf("第%d个方法是:%s:%v\n", i+1, m.Name, m.Type)
	// }

	// v := reflect.ValueOf(s)
	// for i := 0; i < t.NumField(); i++ {
	// 	key := t.Field(i)
	// 	value := v.Field(i).Interface()
	// 	fmt.Printf("%d %s:%s=%v\n", i+1, key.Name, key.Type, value)
	// }

	p := People{Student{ID: 1, Name: "咖啡色的羊驼"}}
	t := reflect.TypeOf(p)
	fmt.Printf("%#v\n", t.Field(0))

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	v := reflect.ValueOf(p)
	fmt.Printf("%#v\n", v.Field(0))

	s := t.FieldByIndex([]int{0, 0})
	fmt.Println(s.Tag)
}

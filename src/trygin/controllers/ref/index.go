package ref

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v, name:%v, kind:%v\n", v, v.Name(), v.Kind())
}

// TypeOf .
func TypeOf(C *gin.Context) {
	var a float32 = 3.14
	reflectType(a)

	var b int64 = 100
	reflectType(b)

	var c rune
	reflectType(c)

	type persion struct {
		name string
		age  int
	}
	var d = persion{name: "人名", age: 18}
	reflectType(d)

	type book struct {
		title string
	}
	var e = book{title: "书名"}
	reflectType(e)

	var f = [...]int{2, 3, 5, 8}
	reflectType(f)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int() 从反射中获取整型的原始值, 然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float() 从反射中获取浮点型的原始值, 然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float() 从反射中获取浮点型的原始值, 然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem() 指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// ValueOf .
func ValueOf(C *gin.Context) {
	var a float32 = 3.14
	var b float64 = 100
	reflectValue(a)
	reflectValue(b)
	c := reflect.ValueOf(10)
	fmt.Printf("type c is :%T\n", c)
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(300)
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 修改的是副本
	if v.Kind() == reflect.Int64 {
		v.SetInt(300)
	}
}

// SetValue 通过反射修改指针对应的值
func SetValue(C *gin.Context) {
	var a int64 = 100
	reflectSetValue(&a) // 传入的不是指针会panic
	fmt.Println(a)
	var b int64 = 100
	reflectSetValue2(&b)
	fmt.Println(b)
}

// NilValid IsNil()判断是否为空指针 IsValid()判断返回值是否有效
func NilValid(C *gin.Context) {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找“abc”字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())

	// map
	c := map[string]int{"哪吒": 1}
	fmt.Println("map中不存在的建:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("哪吒")).IsValid())
}

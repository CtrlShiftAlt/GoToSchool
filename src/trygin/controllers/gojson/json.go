package gojson

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Person 人类
type Person struct {
	Name   string  `json:"name,omitempty"`
	Age    int64   `json:"age,omitempty,string"`
	Weight float64 `json:"weight,omitempty"`
}

// JSONEncode 解析
func JSONEncode(C *gin.Context) {
	p1 := Person{
		Name: "名字",
		Age:  18,
		// Weight: 60,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	jsonstr := `{"name":"名字","age":"18"}`
	var p2 Person
	err = json.Unmarshal([]byte(jsonstr), &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}

// JSONDecode 解析
func JSONDecode(C *gin.Context) {

}

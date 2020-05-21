package testMapStruct

import (
	"fmt"
	"github.com/goinggo/mapstructure"
)

type Person struct {
	N string     `json:"name"`
	A int        `json:"age"`
}

func Map2Struct() {
	mapInstance := make(map[string]interface{})
	mapInstance["name"] = "liang637210"
	mapInstance["age"] = 28

	var person Person
	//将 map 转换为指定的结构体
	err := mapstructure.Decode(mapInstance, &person)
	if  err != nil {
		fmt.Print(err)
	}
	fmt.Print("map2struct后得到的 struct 内容为:%v", person)
}
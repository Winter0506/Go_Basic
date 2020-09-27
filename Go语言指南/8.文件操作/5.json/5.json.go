// json基本介绍
// JSON是一种轻量级的数据交换格式。易于机器解析和生成，并有效提升网络传输效率，
// 通常程序在网络传输时会先将数据（结构体、map等）序列化成json字符串，
// 到接收方得到json字符串时，再反序列化恢复成原来的数据类型（结构体、map等）

package main

import (
	"encoding/json"
	"fmt"
)

// json的序列化
// 结构体序列化
// 对于结构体的序列化，如果我们希望序列化后的key的名字可以自定义，可以加一个tag标签

// Monster ...
type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "红孩儿",
		Age:      200,
		Birthday: "2020-6-28",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	// 将monster反序列化
	data, err := json.Marshal(&monster)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("monster序列化后=%v\n", string(data))
	// monster序列化后={"Name":"红孩儿","Age":200,"Birthday":"2020-6-28","Sal":8000,"Skill":"牛魔拳"}
}

// map序列化
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, err := json.Marshal(a)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

// 切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "南京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("a map 序列化后=%v\n", string(data))

}

// json反序列化是指，将json字符串反序列化成对应数据类型（比如结构体、map、切片）的操作

// 反序列化结构体
func unmarshalStruct() {
	// 直接写的str是需要转义的，如果是程序获得的，不需要转义，内部已经转义
	// 这儿为什么必须写monster_name
	str := "{\"monster_name\":\"红孩儿\",\"Age\":200,\"Birthday\":\"2020-6-28\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)

	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("monster反序列化后 monster=%v monster.Name=%v\n", monster, monster.Name)
}

// 反序列化map
func unmarshalMap(){
	//直接写的str是需要转义的，如果是程序获得的，不需要转义，内部已经转义
	str := "{\"adddress\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	var a map[string]interface{}
	// 注意:反序列化map,不需要make make操作被封装到Unmarshal函数中
	err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("map反序列化后 a=%v \n", a)
	// map反序列化后 a=map[adddress:洪崖洞 age:30 name:红孩儿] 

}

// 反序列化切片
func unmarshalSlice(){
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"

	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v \n", slice)
}

func main() {
	testStruct()
	fmt.Println("------------")
	testMap()
	fmt.Println("------------")
	testSlice()
	fmt.Println("------------")
	unmarshalStruct()
	fmt.Println("-------------")
	unmarshalMap()
	fmt.Println("-------------")
	unmarshalSlice()
}

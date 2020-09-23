package main

import (
	"encoding/json"
	"fmt"
)

// Golang面向对象编程说明
// Golang也支持面向对象编程oop，但是和传统的面向对象编程有区别。
// Golang没有类class，struct和其他语言类class同等地位，基于struct来实现oop特性
// Golang面向对象编程非常简洁，去掉传统oop继承、方法重载、构造函数、析构函数、隐藏this指针等
// Golang任然有面向对象编程的继承、封装、多态,只是实现方式不一样

func f1() {
	// 定义Stu结构体
	/*
		1.结构体是自定义的数据类型，代表一类事物
		2.结构体变量（实例）是具体的，实际的，代表一个具体变量
	*/
	type Stu struct {
		Name    string
		Age     int
		Address string
	}
	// 创建Stu变量
	var stu1 Stu
	stu1.Name = "小王"
	stu1.Age = 12
	stu1.Address = "下沙"
	fmt.Println(stu1.Name)
	fmt.Println(stu1.Age)
	fmt.Println(stu1.Address)
}

func f2() {
	//如果结构体字段类型是：指针、slice、和map的零值是nil，即没有分配空间
	//如果需要使用这样的字段，需要先make，才能使用
	type Person struct {
		Name   string
		Scores [5]float64
		prt    *int
		slice  []int
		map1   map[string]string
	}

	var p1 Person
	if p1.prt == nil && p1.slice == nil && p1.map1 == nil {
		fmt.Println("确实是空")
	}
	// 使用指针 slice map 一定要先make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"

	// 创建结构体变量和访问结构体字段
	// 方式1：直接声明
	// var person Person
	// 方式2：  ！！！！！
	// var person2 Person = Person{}
	// var num int = 9
	// var ptr1 *int = &num
	// person2 = Person{"mary",{20.87,12,134,15,45}, ptr1}  // 这个不会写

	// 方式3：
	var p3 *Person = new(Person)
	(*p3).Name = "smith" // ok
	p3.Name = "join"     // ok
	// 两种方式都可以，go设计者为了程序员使用方便，底层对p3.name进行了处理

	// 方式4：
	var person *Person = &Person{}
	(*person).Name = "smith" // ok
	person.Name = "join"     // ok
}

func f3() {
	// struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见使用场景就是序列化和反序列化
	type Monster struct {
		Name string "json:\"name\""  //json:"name" 就是tag
		Age  int    "json:\"age\""
	}

	monster := Monster{"姜子牙",20}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}

func main() {
	fmt.Println("hello")
	fmt.Println("----------")
	f3()
}

package main

import "fmt"

func f1() {
	// 快速入门
	// map是key-value数据结构
	// 基本语法：var 变量名 map[keytype]valuetype
	// keyType通常为int、string
	// valueType通常为:数字，string、map、struct
	// 例如：var a map[string]map[string]string
	// 声明不会分配内存，初始化需要make
	var a map[string]string
	a = make(map[string]string, 10)
	a["01"] = "宋江"
	a["02"] = "吴用"
	a["01"] = "武松" //key不能重复，如果重复了，以最后一个key-value为准
	// map是无序的
}

func f2() {
	// 使用方式1
	var a map[string]string
	a = make(map[string]string, 10)
	a["01"] = "刘备"
	a["02"] = "关羽"
	// 使用方式2
	cities := make(map[string]string)
	cities["n01"] = "北京"
	cities["n02"] = "天津"
	// 使用方式3
	heroes := map[string]string{
		"hero01": "曹操",
		"hero02": "方腊",
		"hero03": "如来",
	}
	heroes["hero04"] = "王熙凤"
	fmt.Println("heroes=", heroes)
}

func f3() {
	// 存放学生信息,每个学生有name和sex信息
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"

	studentMap["stu02"] = make(map[string]string, 2)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
}

func f4() {
	// map增加和更新：map["key"] = value //如果key还没有，就是增加，如果可以存在就是修改
	// map删除
	// delete(map,"key"),delete是一个内置函数，如果key存在就删除key-value，如果key不存在，不操作
	cities := make(map[string]string)
	cities["n01"] = "北京"
	cities["n02"] = "天津"
	delete(cities, "no1")
	// 如果要删除map的所有key，没有专门的一个方法，需哟遍历一下key 或者 map = make(..) make一个新的，让原来的成为垃圾，被GC回收
	cities = make(map[string]string)
}

func f5() {
	cities := make(map[string]string)
	cities["n01"] = "北京"
	cities["n02"] = "天津"
	val, ok := cities["n01"]
	if ok {
		fmt.Printf("有n01 key 值为%v \n", val)
	} else {
		fmt.Println("没有n01 key")
	}
}

func main() {
	fmt.Println("Hello")
	// f2()
	// f3()
	f5()
}

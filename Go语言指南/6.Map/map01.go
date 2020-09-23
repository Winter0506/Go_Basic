package main

import (
	"fmt";
	"sort"
)

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

func f6(){
	cities := make(map[string]string)
	cities["n01"] = "北京"
	cities["n02"] = "天津"
	cities["n03"] = "上海"

	for k,v := range cities{
		fmt.Printf("k=%v v=%v\n",k,v)
	}

	fmt.Println("-----------------")

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string,2)
	studentMap["stu01"]["name"]="tom"
	studentMap["stu01"]["sex"]="男"
	studentMap["stu02"] = make(map[string]string,2)
	studentMap["stu02"]["name"]="mary"
	studentMap["stu02"]["sex"]="女"

	for _,v1 := range studentMap{
		for k2,v2 := range v1{
			fmt.Printf("k2=%v,v2=%v",k2,v2)
		}
		fmt.Println()
	}


}

func f7(){
	// map的长度
	cities := make(map[string]string)
	cities["n01"] = "北京"
	cities["n02"] = "天津"
	cities["n03"] = "上海"
	fmt.Println(len(cities))   // 求长度
}

func f8(){
	// map切片，map数量就可以动态变化
	var monsters []map[string]string
	monsters = make([]map[string]string,2)
	if monsters[0] == nil{
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil{
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "黄鼠狼"
		monsters[1]["age"] = "400"
	}
	// 通道append函数，动态增加monster
	newMonster := map[string]string{
		"name":"大司马",
		"age":"35",
	}
	monsters = append(monsters,newMonster)
	fmt.Println(monsters)
}

func f9()  {
	// golang没有专门的方法针对map的key进行排序；
	// map是无序的，每次遍历得到的输出都不一样，不是按照添加的顺序存放
	// golang中map的排序，是先将key 进行排序。然后根据key值遍历输出
	
	map1 := make(map[int]int,10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	//先将key值放入切片
	//对切片排序
	//遍历切片，按key来输出map的值
	var keys []int
	for k := range map1{
		keys = append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys{
		fmt.Printf("map[%v]=%v \n",k,map1[k])
	}
}

func f10(){
	// map是引用类型，遵守引用类型传递的机制，在一个函数接收map后，会直接修改原来的map
	// map的容量达到后，在想map增加元素，会自动扩容，并不会发生panic，map可以动态的增长
	// map的value经常使用struct类型，更适合管理复杂的数据
	type Stu struct{
		Name string
		Age int
		Address string
	}
	// 创建Stu变量
	var stu1 Stu
	stu1.Name = "小王"
	stu1.Age = 12
	stu1.Address = "下沙"
	students := make(map[string]Stu,10)
	students["no1"] = stu1

	stu2 := Stu{"小李",13,"上海"}

	students["no2"] = stu2 

	for k,v := range students{
		fmt.Printf("学生编号%v \n",k)
		fmt.Printf("学生姓名%v \n",v.Name)
		fmt.Printf("学生年龄%v \n",v.Age)
		fmt.Printf("学生城市%v \n",v.Address)
	}

}

func main() {
	fmt.Println("Hello")
	// f2()
	// f3()
	// f5()
	// f6()
	// f7()
	// f8()
	// f9()
	f10()

}

// 类型断言
//如何将一个接口变量，赋给自定义类型的变量 --->断言

package main

import (
	"fmt"
)

type Point struct{
	x int
	y int
}

func TypeJudge(items... interface{}){
	for index,v := range items{
		switch v.(type){
		case bool:
			fmt.Printf("第%v个参数，是bool类型，值为：%v\n",index,v)
		case float32:
			fmt.Printf("第%v个参数，是float32类型，值为：%v\n",index,v)
		case int, int32 ,int64:
			fmt.Printf("第%v个参数，是整数类型，值为：%v\n",index,v)						
		
		case string:
			fmt.Printf("第%v个参数，是string类型，值为：%v\n",index,v)	
		default:
			fmt.Printf("第%v个参数，类型不确定，值为：%v\n",index,v)	
		}
	}
}

func main(){
	var a interface{}
	var p Point = Point{1,2}
	a = p
	fmt.Println(a) //{1 2}
	var p2 Point
	p2 = a.(Point) // 断言
	fmt.Println(p2) // {1 2}

	fmt.Println("------------")

	// 断言时加上检测机制
	var x interface{}

	var b2 float32 = 2.1
	x = b2
	if y, ok := x.(float32); ok{
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T值是=%v",y,y)
	}else{
		fmt.Println("vonvert fail")
	}

	fmt.Println("--------------")

	var n1 float32 = 1.1
	var n2 int = 30
	var name string = "tom"
	address := "北京"

	TypeJudge(n1,n2,name,address)
}
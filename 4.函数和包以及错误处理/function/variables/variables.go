package main
import "fmt"

// 1.函数内部声明/定义的变量叫局部变量,作用域仅限于函数内部
func test(){
	// age和name的作用域只在test函数内部
	age := 10
	name := "tom"
	fmt.Println("age=",age)
	fmt.Println("name=",name)
}


// 2.函数外部声明/定义的变量叫做全局变量,作用域在整个包都有效,如果其首字母大写,则作用域在整个程序有效
var age int = 50 //整个包有效
// Name ...
var Name string = "jack" // 首字母大写，整个程序有效

//3.如果变量是一个代码块,比如for/if中,那么这个变量的作用域就在该代码块

func main(){
	fmt.Println("Hello")
	test()
	fmt.Println("------------")
	fmt.Println("age=",age) 
	fmt.Println("Name=",Name)
}
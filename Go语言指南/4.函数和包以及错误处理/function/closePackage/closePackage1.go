package main

import (
	"fmt"
	"strings"
)

//AddUpper 累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

// AddUpperStr haha
func AddUpperStr() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += "!"
		fmt.Println("str=", str)
		return n
	}
}

// 闭包的最佳实践
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) // 11  //n是初始化一次，因此每调用一次就进行累计
	fmt.Println(f(2)) // 13  //11+2
	fmt.Println(f(3)) // 16	 //13+3

	f1 := AddUpperStr()
	fmt.Println(f1(1))
	fmt.Println(f1(2))
	fmt.Println(f1(3))
	/*
		str= hello!
		11
		str= hello!!
		13
		str= hello!!!
		16
	*/
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))   //winter.jpg
	fmt.Println("文件名处理后=", f2("bird.jpg")) //bird.jpg
	fmt.Println("文件名处理后=", f2("winter.jpg"))   
	fmt.Println("文件名处理后=", f2("bird"))   



	// 1.返回的匿名函数和makeSuffix(suffix string)的suffix变量组合成一个闭包，因为返回的函数引用到suffix这个变量
	// 2.我们体会一下闭包的好处，闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复使用。
}

// 总结:
/*
	1.AddUpper是一个函数，返回的数据类型是fun(int)int   注意 这个地方是重点
	2.闭包的说明：返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包。
	3.可以这样理解:闭包是类，函数是操作，n是字段。函数和它使用的n构成闭包
	4.当我们反复的调用f函数时，因为n是初始化一次，因此每调用一次就进行累计
	5.我们要搞清楚闭包的关键，就是要分析出返回的函数它使用到哪些变量，因为函数和它引用到的变量共同构成闭包
	6.对上面代码的一个更改，加深对闭包的理解
*/

package main

import "fmt"

// 匿名函数
// Go支持匿名函数，匿名函数就是没有名字的函数，如果我们某个函数只是希望使用一次，可以考虑匿名函数，
// 匿名函数也可以实现多次调用。

// 方式三：全局匿名函数：如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以再程序有效
var (
	Fun1 = func (n1 int, n2 int)  int{
		return n1*n2
	}
)

func main() {
	//方式一：在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res1 := func(n1 int, n2 int) int{
		return n1+n2
	}(10,20)
	fmt.Println("res1=",res1)

	//方式二：将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数
	a := func(n1 int, n2 int) int{
		return n1 - n2
	}
	res2 := a(10,30)
	fmt.Println("res2=",res2)

	res3 := Fun1(10,20)
	fmt.Println("res3=",res3)

}


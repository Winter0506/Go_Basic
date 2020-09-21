// 数组可以存放多个同一类型数据。数组也是一种数据类型，在Go中，数组也是值类型。
package main

import "fmt"

func main() {
	/*
		//定义
		var 数组名 [数组大小]数据类型
		var a [5]int
		赋值 a[0]=1
	*/
	var intArr [3]int
	fmt.Println(intArr) //各个元素有默认值0
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	/*
		[0 0 0]
		[10 20 30]
	*/

}

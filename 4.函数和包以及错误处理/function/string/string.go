package main

import (
	//"strings";
	"fmt"
	"strconv"
)

func main() {
	// 1.统计字符串的长度，按照字节len(str)
	str := "hello北京"
	fmt.Println("str len = ", len(str)) // 1

	fmt.Println("----------")

	// 2.字符串遍历,同时处理有中文的问题 r:=[]rune(str)
	str2 := "hello上海"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	fmt.Println("----------")

	// 3.字符串转整数:   这儿有点儿乱哈！！！！
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
		fmt.Println("转成的结果是", n)
	} else {
		fmt.Println("转成的结果是", n)
	}
	n1, err1 := strconv.Atoi("100")
	if err != nil {
		fmt.Println("转换错误", err1)
		fmt.Println("转成的结果是", n1)
	} else {
		fmt.Println("转成的结果是", n1)
	}
	fmt.Println("----------")
	/*
		转换错误 strconv.Atoi: parsing "hello": invalid syntax
		转成的结果是 0
		转换错误 <nil>
		转成的结果是 100
	*/

	// 4.整数转字符串
	str3 := strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T", str3, str3)
	fmt.Println("-----------")

}

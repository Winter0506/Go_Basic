package main

import "fmt"

func main()  {
	fmt.Println("haha")

	// 1.string底层是一个byte数组，因此string也可以进行切片处理
	str := "helloworld"
	slice := str[5:]
	fmt.Println("slice=",slice)

	// 2.string是不可变的，也就说不能通过str[0] = 'z'方式修改字符串
	// 3.如果需要修改字符串，可以string -》[]byte/或者[]rune ->修改 -》重写转成string
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=",str)
	//转成[]byte后，可以处理英文和数字，但不能处理中文，中文占3个字节
	//解决方法是转成[]rune即可，[]rune是按字符处理，兼容汉字
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str=",str)
}
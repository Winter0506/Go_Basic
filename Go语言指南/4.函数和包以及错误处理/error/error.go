// 错误处理
// Go语言追求简洁优雅，所以，Go语言不支持传统的try。。。catch。。。finally这种处理
// Go中引入的处理方式：defer、panic、recover
// 简单描述:Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

package main

import (
	"errors"
	"fmt"
)

// 1.默认情况下，当发生错误后（panic）,程序就会退出（崩溃）
func test() {
	// 使用defer + recover来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// 2.自定义错误
// Go程序中，也支持自定义错误，使用errors.New 和 panic内置函数

// errors.New("错误说明")，会返回一个error类型的值，表示一个错误
// panic内置函数，接收一个interface{}类型的值作为参数。可以接收error类型的变量，输出错误信息，并退出程序。
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取..
		return nil
	}
	return errors.New("读取文件错误")
}

func test02() {
	err := readConf("config2.ini")
	if err != nil {
		fmt.Println("进入了...")
		// 如果读取文件发生错误,就输出这个错误，并终止程序
		panic(err)
	}
	// 下面这句话不会执行
	fmt.Println("test02()继续执行...")
}

func main() {
	test()
	test02()
	fmt.Println("---------------")

	fmt.Println("main()下面的代码")
}

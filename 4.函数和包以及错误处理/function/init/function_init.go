package main

import "fmt"

// 每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，也就是说init会在mian函数前被调用。

// init函数，通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...")
}

// init函数的注意事项和细节

// 如果一个文件同时包含全局变量定义，init函数和mian函数，则执行顺序：全局变量定义-> init-> main
// init函数最主要的作用，就是完成一些初始化的工作
// 细节说明：如果main.go 和utils.go都包含变量定义、init函数，流程为：utils：变量定义-> init-> main：变量定义-> init-> main
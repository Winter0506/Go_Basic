package main

import "fmt"

// 在Go中，函数分为：自定义函数、系统函数
// 函数的基本语法
// func 函数名（形参列表）（返回值列表）{
//    执行语句..
//    return 返回值列表
// }

// 函数实例
func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误")
	}
	return res
}

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'
	result := cal(n1, n2, operator) //调用函数
	fmt.Println("result=", result)
}

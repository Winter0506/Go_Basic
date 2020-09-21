// 执行一个函数时，就创建一个新的受保护的独立空间
// 函数的局部变量是独立的，不会相互影响
// 递归必须向退出递归的条件逼近，否则就是无限递归
// 当一个函数执行完毕，或者遇到return，就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或者返回时，该函数本身也会被系统销毁

package main

import "fmt"

// Factorial 阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// 斐波那契数列
func fibonacci(n uint64) (result uint64) {
	if n > 2 {
		return fibonacci(n-2) + fibonacci(n-1)
	}
	return n
}

func main() {
	var i int = 3
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
	fmt.Printf("%d 的斐波那契数列是%d\n", i, fibonacci(uint64(i)))
}

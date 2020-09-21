package main

/*
cannot find package "add" in any of:
	D:\Go\GOROOT\src\add (from $GOROOT)
	D:\Go\GOPATH\src\add (from $GOPATH)
*/
import (
	"fmt"
	// "add"   所以必须在上面两个文件夹下才能导入?
)

// 包的本质实际上就是创建不同的文件夹，来存放程序文件；
// go的每一个文件都是属于一个包的，也就是说go是以包的形式来管理文件和项目目录结构的

// 包的作用
// 区分相同名字的函数、变量等标识符
// 当程序文件很多时，可以很好的管理项目
// 控制函数、变量等访问范围，即作用域

// 包的相关说明
// 打包基本语法：package 包名
// 引入包的基本语法：import "包的路径"

// 当一个文件要使用其他包函数或者变量时，需要引入对应的包
// 方式一
// import "包名"
// 方式二
// import (
//   	"包名"
//   	"包名"
// )

// 为了让其他包的文件，可以访问到本包的函数，则该函数名的首字母需要大写，类似其他语言的public

// 在访问其他包函数、变量，其语法是包名.函数名

// 如果包名较长，Go支持给包取别名，注意细节：取别名后，原来的包名就不能使用了

/*
package main
import(
	"fmt"
	util "go_code/utils"
)
*/

// 在同一包下，不能有相同的函数名（也不能有相同的全局变量名）

// 如果你要编译成一个可执行程序文件，就需要将这个包声明为main，即package main。如果你是写一个库，包名可以自定义


// package指令在文件第一行，然后是import指令
// 在import包时，路径从$GOPATH的src下开始，不用带src，编译器会自动从src下开始引入

// return 语句
/*
基本语法和说明
Go函数支持返回多个值，这一点是其他编程语言没有的
func 函数名(形参列表)(返回值类型列表){
    语句
    return 返回值列表
}
1.如果返回多个值，在接收时，希望忽略某个返回值，则使用_符号表示占位忽略
2.如果返回值只有一个，返回值类型列表可以不写()
*/

// 返回一个值
func f1(n1 int, n2 int)(int){
	var res int
	res = n1 + n2
	return res
}
// 返回一个值返回值类型省略()
func f2(n1 int, n2 int) int{
	var res int
	res = n1 + n2
	return res
}

// 返回多个值
func f3(n1 int, n2 int) (int,int){
	var sum int
	var sub int
	sum = n1 + n2
	sub = n2 - n1
	return sum, sub
}

func main(){
	var n1 int = 10
	var n2 int = 20
	fmt.Println("返回一个值:")
	sum1 := f1(n1,n2)
	fmt.Println("sum1的值为:",sum1)
	fmt.Println("返回一个值返回值类型省略():")
	sum2 := f2(n1,n2)
	fmt.Println("sum2的值为:",sum2)
	fmt.Println("返回多个值:")
	sum3, sub1 := f3(n1,n2)
	fmt.Println("sum3,sub1的值为:",sum3,sub1) 
	sum4, _ := f3(n1,n2)
	fmt.Println("省略一个值的另一个值为:",sum4)
}
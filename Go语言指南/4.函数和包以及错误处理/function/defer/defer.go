// 在函数中，程序员经常需要创建资源（比如：数据库连接、文件句柄、锁等），
// 为了在函数执行完毕后，及时的释放资源，Go的设计者提供defer（延时机制）

package main

import "fmt"

func sum(n1 int, n2 int) int {
	//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1) // 第三个输出
	defer fmt.Println("ok2 n2=", n2) // 第二个输出

	res := n1 + n2
	fmt.Println("ok3 res=", res) // 第一个输出
	return res
}

// defer的注意事项和细节
/*
当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中，然后继续执行函数下一个语句
当函数执行完毕后，再从defer栈中，依次从栈顶取出语句执行（遵守栈 先入后出的机制）
在defer将语句放入到栈时，也会将相关的值拷贝同时入栈。
*/

func sum1(n1 int, n2 int) int {
	//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1) //3  此处压入栈的n1 为10
	defer fmt.Println("ok2 n2=", n2) //2 此处压入栈的n2 为20
	//改变下n1 n2的值
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res) //1
	return res
}

/*
在golang编程中的通常做法，创建资源后（打开文件、获取数据库连接、锁资源），可以执行defer file.Close() defer connect.Close()
在defer后，可以继续使用创建资源
当函数完毕后，系统会依次从defer栈中，取出语句，关闭资源
这种机制，非常简洁，程序员不用再为再什么时候关闭资源而烦恼
*/

// defer的最佳实践
/*
func test(){
	file = openfile(文件名)
	defer file.close()//关闭文件资源
	//其他代码
}
func test(){
    //释放数据库资源
    connect = openDatabase()
    defer connect.close()
    //其他代码
}
*/

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
	/*
		ok3 res= 30
		ok2 n2= 20
		ok1 n1= 10
		res= 30
	*/
	res1 := sum1(10, 20)
	fmt.Println("res1=", res1) //4
	/*
		ok3 res= 32
		ok2 n2= 20
		ok1 n1= 10
		res1= 32
	*/
}

package main

import "fmt"


// 算术运算符
// 赋值运算符
// 比较运算符/关系运算符
// 逻辑运算符
// 位运算符
// 其他运算符

func f1()  {
	// 算术运算符
	// 对于除号"/"，它的整数除和小数除是有区别的：整数之间做除法，只保留整数部分而舍弃小数部分。

	// 当对个数取模时，可以等价a%b = a-a/b*b

	// Golang的自增自减只能当作一个独立语言使用时，不能如下使用
	// var i int = 8
	// a = i++ //错误
	// a = i-- //错误
	// if i++ > 0 {//错误
	
	// }

	// Goalng的++和--只能写在变量的后面，不能写在变量的前面。
	// var i int =  1
	// i++ //对
    // ++i //错误
    // i-- //对
    // --i //错误
	
}

func f2(){
	// 逻辑运算符
	// && 也叫短路与：如果第一个条件为false，则第二个条件不会判断，最终结果为false
	// || 也叫短路或：如果第一个条件为true，则第二个条件不会判断，最终结果为true
}

func f3(){
	// 赋值运算符
}

func f4(){
	// 位运算符
	// & 按位与运算符”&“是双目运算符。其功能是参与运算的两位数对应的二进制位相与。运算按规则是：同时为1，结果为1，否在为0
	// | 按位或运算符”|“是双目运算符，其功能是参与运算的两位数各对应的二进位相或。运算规则是：有一个为1，结果为1，否在为0
	// ^ 按位异或运算符”^“是双目运算符。其功能是参与运算的两位数各对应的二进位相异或。运算规则是：当二进位不同时，结果为1，否在为0
	// << 左移运算符”<<“是双目运算符。其功能把”<<“左边的运算数的各二进位全部左移若干位，高位丢弃，低位补0.左移n位就是乘以2的n次方
	// >> 右移运算符">>"是双目运算符。其功能是把”>>“左边的运算数的各二进位全部右移若干位，右移n位就是除以2的n次方
}

func f5(){
	// 其他运算符
	// &  返回变量存储地址  &a；将给出变量的实际地址
	// *  指针变量        *a；是一个指针变量
}

func f6(){
	// Go语言明确不支持三元运算符
	var n int
	var i int = 10
	var j int = 12
	// 三元运算符
	//n = i>j?i:j
	if i > j{
		n = i
	} else{
		n = j
	}
	fmt.Printf("n=%d",n)
}

func f7()  {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名 ")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄 ")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水 ")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试 ")
	fmt.Scanln(&isPass)
	fmt.Printf("名字是%v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n",name,age,sal,isPass)

	fmt.Println("请输入您的姓名,年龄,薪水,是否通过考试,使用空格隔开")
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
}

func f8(){
	var i int = 5
	fmt.Printf("%b \n",i) //二进制输出
	var j int = 011 //八进制
	fmt.Println("j",j)
	var k int = 0x11 //16进制
	fmt.Println("k",k)
}

func main(){
	// f6()
	// f7()
	f8()
}
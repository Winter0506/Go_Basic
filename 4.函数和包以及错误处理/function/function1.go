package main

import "fmt"

/*
函数的形参列表可以是多个，返回值列表也可以是多个

形参列表和返回值列表的数据类型可以是值类型和引用类型

函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写类似于public，首之母小写，只能被本包文件使用，类似于private

函数中的变量是局部的，函数外不生效

基本数据类型和数组默认都是值传递，即进行值拷贝。在函数内修改，不会影响到原来的值

*/
func test01(n1 *int) int{
	// 如果希望函数内的变量能够修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量
	*n1 = *n1 + 10
	return *n1
}

//07.Go函数不支持函数重载
// func test02(n1 int){}
// func test02(n1 int,n2 int){}//error 不支持重载


// 在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
func getSum(n1 int, n2 int) int{
	return n1+n2
}

// 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int{
	return funvar(num1, num2)
}

// 为了简化数据类型定义，Go支持自定义数据类型
// 基本语法： type 自定义数据类型名 数据类型 //相当于一个别名
// 案例：type myInt int //这时myInt就等价int来使用
//      type mySum func(int,int)int //这时mySum就等价一个函数类型 func(int,int) int 

type myFunType func(int, int) int

func myFun2(funvar myFunType, num1 int, num2 int) int{
	return funvar(num1, num2)
}

// 支持对函数返回值命名
/*
func getSumAndSub(n1 int,n2 int)(sum int,sub int){
    sub = n1 - n2
    sum = n1 + n2
    return   // 这个地方只写了一个return也可以
}

a1,b1 := getSumAndSub(1,2)
fmt.Printf("a=%v b=%v",a1,b1)
*/
func getSumAndSub(n1 int,n2 int)(sum int,sub int){
    sub = n1 - n2
    sum = n1 + n2
    return   // 这个地方只写了一个return也可以
}

// 使用_标识符，忽略返回值
/*
func cal(n1 int,n2 int)(sum int,sub int){
    sum = n1 + n2
    sub = n1 - n2
    return
}

func main(){
    res1,_ = cal(10,20)
    fmt.Printf("res1=%d",res1)
}
*/


// Go支持可变参数
// 支持0到多个参数
// func sum(args... int) sum int{}
// 支持1到多个参数
// func sum(n1 int,args... int) sum int{}

/*
说明：
1.args是slice切片，通过args[index]可以访问到各个值。
2.如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表后面
案例：写一个函数sum，可以求出1到多个int的和
*/
func sum(n1 int, args... int) int{
	sum := n1
	fmt.Println(sum)
	for i:=0; i< len(args); i++{
		sum+=args[i]
	}
	fmt.Println(sum)
	return sum
}

func main()  {
	fmt.Println("hello")
	// 这个指针的函数终于运行通过了
	value := 10
	var n1 *int = &value
	result01 := test01(n1)
	fmt.Println(result01)

	fmt.Println("-------------------")

	a := getSum
	fmt.Printf("a的类型是%T，getSum的类型是%T\n",a,getSum)

	res := a(10,20) //等价于 res:=get(10,20)
	fmt.Println("res=",res)

	fmt.Println("-------------------")
	res2 := myFun(getSum,50,60)
	fmt.Println("res2=",res2)

	fmt.Println("-------------------")
	res3 := myFun2(getSum,100,200)
	fmt.Println("res3:=",res3)

	fmt.Println("-------------------")
	a1,b1 := getSumAndSub(1,2)
	fmt.Printf("a=%v b=%v\n",a1,b1)

	fmt.Println("-------------------")
	res4 := sum(10,0,-1,90,10,100)
	fmt.Println("res4=",res4)

}
// 数组可以存放多个同一类型数据。数组也是一种数据类型，在Go中，数组也是值类型。
package main

import "fmt"

func test(arr *[3]int)  {
	(*arr)[0]=88
}

func modify(arr [3]int){
	fmt.Println(arr)
}

func main() {
	/*
		//定义
		var 数组名 [数组大小]数据类型
		var a [5]int
		赋值 a[0]=1
	*/
	var intArr [3]int
	fmt.Println(intArr) //各个元素有默认值0
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	/*
		[0 0 0]
		[10 20 30]
	*/
	fmt.Printf("intArr的地址=%p intArr[0]的地址=%p intArr[1]的地址=%p intArr[2]的地址=%p\n",&intArr,&intArr[0],&intArr[1],&intArr[2])
	/*
	1.数组的地址可以通过数组名来获取&intArr
	2.数组的第一个元素的地址，就是数组的首地址
	3.数组的各个元素的地址间隔是依据数组的类型决定，比如int64 ->8 int32->4
	*/
	// intArr的地址=0xc00000a460 intArr[0]的地址=0xc00000a460 intArr[1]的地址=0xc00000a468 intArr[2]的地址=0xc00000a470
	// 输出结果里面为什么不是相隔8个字节

	// 1.数组的使用
	// 从终端循环输入5个价格，保存到float64数组，并输出
	/*
	var price [5]float64
	for i :=0; i<5; i++ {
		fmt.Printf("请输入第%v个数:",i+1)
		fmt.Scanln(&price[i])
	}
	fmt.Println(price)
	*/

	// 四种初始化数组方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	var numArr02 = [3]int{5,6,7}
	var numArr03 = [...]int{8,9,10}
	var numArr04 = [...]int{1:800,0:900,2:999} //指定下标
	fmt.Println(numArr01)
	fmt.Println(numArr02)
	fmt.Println(numArr03)
	fmt.Println(numArr04)
	// 类型推导
	strArr05 := [...]string{1:"tom",0:"jack",2:"mary"}
	fmt.Println(strArr05)

	// 2.数组的遍历
	//方式一 常规遍历
	//方式二 for-range结构遍历
	// for index，value :=range array01{

	// }
	// 1.index数组的下标
	// 2.value该下标对应的值
	// 3.他们都是for循环内可见的局部变量
	// 4.如果不想使用下标index，可以替换为"_"
	// 5.index和value的名称不是固定的。可以自己改变

	heroes := [...]string{"宋江","吴用","卢俊义"}

	for _, v := range heroes {
		fmt.Printf("元素的值=%v\n",v)
	}

	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n",i,v)
		fmt.Printf("heroes[%v] = %v\n",i,heroes[i])
	}

	// 数组使用的注意事项和细节
	/*
	数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化

	var arr []int 这时arr就是一个slice切片，切片后面讲解

	数组中的元素可以是任何数据类型，包括值类型和引用类型，但不能混用

	数组创建后，如果没有赋值，有默认值(零值)

	使用数组的步骤：1.声明数组并开辟空间 2.给数组各个元素赋值（默认零值）3.使用数组

	数组的下标是从0开始的

	数组下标必须在指定范围内使用，否则报panic：数组越界

	Go的数组属值类型，默认情况下值传递，因此会进行值拷贝。数据不相互影响

	如果想在其他函数中，去修改原来的数组，可以使用引用传递（指针方式）
	*/

	test(&numArr01)
	fmt.Println(numArr01)
	fmt.Println("------------")
	// 长度是数组类型的一部分，在传递函数参数时，需要考虑数组长度

	modify(numArr01)

}

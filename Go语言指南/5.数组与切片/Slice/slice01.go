// 先看一个需求：我们需要一个数组用于保存学生成绩，但是学生的个数是不确定的，请问怎么办？解决方案：使用切片
// 切片基本介绍
/*
切片英文是slice

切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制

切片的使用和数组类似，遍历切片、访问切片的元素和求求切片长度len(slice)都一样

切片的长度是可以变化的，因此切片是一个可以动态变化数组

切片定义的基本语法
var 切片名 []类型
var a []int
*/

package main

import "fmt"

func test(slice []int) {
	slice[1] = 200
}

func main() {
	fmt.Println("hello")
	// 2.切片快速入门
	// 方式一：定义一个切片，然后让切片去引用一个已经创建好的数组，比如前面的案例就是这样
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}
	slice := intArr[1:3] //声明一个切片 引用到intArr这个数组 起始下标为1，最后的下标为3（不包含3）
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 = ", slice)
	fmt.Println("slice 的元素个数 = ", len(slice)) // 2
	fmt.Println("slice的容量 = ", cap(slice))    // 切片的容量可以动态变化 目前是4

	// 方式二: 通过make来创建切片
	/*
		基本语法： var 切片名 []type = make([]type,len,[cap])
		参数说明：type数据类型 len：大小 cap：指定切片容量 可选，如果分配cap>=len
	*/
	var slice01 []float64 = make([]float64, 5, 10)
	slice01[1] = 10
	slice01[2] = 20
	fmt.Println("slice 的元素是 =", slice01)
	fmt.Println("slice 的元素个数是 = ", len(slice01))
	fmt.Println("slice 的容量 = ", cap(slice01))
	fmt.Println("----------------")
	/*
		方式一和方式二区别：
		方式一是直接引用数组，这个数组是事先存在的，程序是可见的
		方式二是通过make来创建切片，make也会创建一个数组，是由切片在底层进行维护，程序员是看不见的，
	*/

	//5.切片的遍历
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice02 := arr[1:4] //20,30,40
	// 方式1 循环遍历
	for i := 0; i < len(slice02); i++ {
		fmt.Printf("slice02[%v]=%v\n", i, slice02[i])
	}

	// 方式2 for-range 结构遍历切片
	for index, value := range slice {
		fmt.Printf("index=%v value=%v\n", index, value)
	}
	fmt.Println("---------------")

	// 6. 切片的使用的注意事项和细节讨论
	// 切片初始化时 var slice = arr[startIndex:endIndex] ,从arr数组下标为startIndex取到下标为endIndex-1的元素
	// 切片初始化时，任然不能越界，范围在[0-len[arr]]之间，但是可以动态增长
	/*
		var slice = arr[0:end] 可用简写 var slice = arr[:end]
		var slice = arr[start:len(arr)] 可用简写 var slice = arr[start:]
		var slice = arr[0:len(arr)] 可以简写 var slice = arr[:]
	*/

	/*
		cap是一个内置函数，用于统计切片的容量，即最大可用存放多少个元素。
		切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用
		切片可用继续切片
	*/
	var arr01 [5]int = [...]int{10, 20, 30, 40, 50}
	slice03 := arr01[1:4]
	slice031 := slice[1:2]
	slice031[0] = 100 // 因为arr01,slice03和slice031指向的数据空间是同一个
	fmt.Println("slice031=", slice031)
	fmt.Println("slice03=", slice03)
	fmt.Println("arr=", arr)
	fmt.Println("----------------")

	// 用append内置函数,可以对切片进行动态追加
	var slice04 []int = []int{100, 200, 300}
	slice04 = append(slice04, 400, 500, 600)
	fmt.Println("slice04", slice04)
	slice04 = append(slice04, slice04...)
	fmt.Println("slice04", slice04)

	/*
		切片append操作的底层原理分析：

		切片append操作的本质就是对数组扩容

		go底层会创建一个新的数组newArr（按照扩容后大小）

		将slice原来包含的元素拷贝到新的数组newArr

		slice重新引用到newArr

		newArr是在底层来维护的，程序员不可见
	*/
	// 7.切片的拷贝操作：使用copy内置函数完成拷贝
	var slice05 []int = []int{1, 2, 3, 4, 5}
	var slice06 = make([]int, 10) // 10代表大小
	copy(slice05, slice06)
	fmt.Println("slice05=", slice05)
	fmt.Println("slice06=", slice06)
	// 说明：
	// 1.copy(para1,para2)参数的数据类型是切片
	// 2.slice4和slice5的数据空间是独立，互相不影响，slice4[0]=999 slice[0]仍然是1

	// 8.关于拷贝的注意事项
	var a []int = []int{1, 2, 3, 4, 5}
	var slice07 = make([]int, 1)
	fmt.Println(slice07) // [0]
	copy(slice07, a)
	fmt.Println(slice07) // [1]

	fmt.Println("-------------------")

	// 9.切片是引用类型，所以在传递时，遵守引用传递机制
	var slice08 []int
	var arr02 [5]int = [...]int{1, 2, 3, 4, 5}
	slice08 = arr02[:]
	var slice09 = slice08
	slice09[0] = 100
	fmt.Println("slice09:", slice09)
	fmt.Println("slice08:", slice08)
	fmt.Println("arr02:", arr02)

	fmt.Println("--------------------")
	test(slice09)
	fmt.Println("slice=", slice09)
	fmt.Println("--------------------")

}

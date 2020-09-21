package main

import (
	"fmt"
	// 1.时间和日期相关函数，需要导入time包
	"time"
)

func main() {
	fmt.Println("hello")

	// 2.time.Time 类型，用于表示时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T", now, now)
	fmt.Println("-----------------")

	// 3.如何获取到其他的日期信息
	// 通过now可以获取年月日,时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
	/*
		月=September
		月=9
	*/
	fmt.Println("-----------------")

	// 4.格式化日期时间
	// 方式一：使用Printf或者Sprintf
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v", dateStr)
	fmt.Println("-----------------")

	// 方式二：使用time.Format()方法完成
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	fmt.Println("-----------------")

	// 5.时间常量
	//作用：获取指定时间单位时间，比如100毫秒 100*time.Millisecond
	// 为啥子有错
	/*
		const(
			Nanosecond Duration = 1 //纳秒
			Microsecond = 1000*Nanosecond //微秒
			Millisecond = 1000*Microsecond //毫秒
			Second = 1000*Millisecond //秒
			Minute = 60*Second //分钟
			Hour = 60*Minute //小时
		)
	*/

	//6.结合sleep使用时间常量
	//每个1秒打印一个数字，打印到5时就退出
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
		if i == 5 {
			break
		}
	}
	fmt.Println("---已经结束---")

	// 7.time的Unix和UnixNano的方法
	fmt.Printf("Unix时间戳=%v unixnano时间戳=%v", now.Unix(), now.UnixNano()) //Unix时间戳=1600659510 unixnano时间戳=1600659510776593200
	fmt.Println("--------------")

}

// Golang设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称为Go的内置函数。

// 1.len：用来求长度，比如：string、array、slice、map、channel

// 2.new：用来分配内存，主要用来分配值类型，比如int、float32，struct...返回的是指针
	// num1 := 100 // 内存中 num1 --> 100
	// num2 := new(int)
	// *num2 = 100 //内存中 num2-->0xc04204c098-->100
// 3.make：用来分配内存，主要用来分配引用类型，比如channel、map、slice

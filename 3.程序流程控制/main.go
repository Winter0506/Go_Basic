package main

import "fmt"

//顺序控制 分支控制 循环控制

//分支控制
//单分支控制

// if 条件表达式{
// 	执行代码块
// }

//双分支控制

// if 条件表达式{
// 	执行代码块1
// }else{
// 	执行代码块2
// }

//多分支控制

// if 条件表达式1{
// 	执行代码块1
// }else if条件表达式2{
//  	执行代码块2
// }
// .....
// else{
// 	执行代码块n
// }

//switch分支控制

// switch语句用于基于不同条件执行不同动作，每一个case分支是唯一的，从上到下逐一测试，直到匹配为止。
// 匹配项后面也不需要再加break

// switch 表达式{
// 	case 表达式1，表达式2，..:
// 		语句块1
// 		fallthrough//穿透 如果有这个则会继续执行下面的case
// 	case 表达式3，表达式4，..:
// 		语句块2
// 	...
// 	default://没有任何case匹配 折执行default 不是必须的
// 		语句块
// }

// switch的使用的注意事项和细节
func f1() {
	// case/switch 后是一个表达式（即：常量值、变量、一个返回值的函数等）

	// case后的各个表达式的值的数据类型，必须和switch的表达式数据类型一致

	// case 后面可以带多个表达式，使用逗号间隔。
	var n1 int32 = 5
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5:
		fmt.Println("ok1")
	// case 5: //error 前面有常量5了
	//	fmt.Println("ok2")
	case 6:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}
}

func f2() {
	// case 后面的表达式如果是常量值，则要求不能重复

	// case后面不需要带break，执行case会退出switch，如果没有匹配到则执行default

	// default语句不是必须的

	// switch后可以不带表达式，类似 if --else 分支来使用
	var age int = 10
	switch {
	case age == 1:
		fmt.Println("age==1")
	case age == 2:
		fmt.Println("age==2")
	default:
		fmt.Println("没有找到匹配")
	}

	// case中可以对范围进行判断
	var score int = 90
	switch {
	case score > 90:
		fmt.Println("成绩优秀")
	case score >= 70 && score <= 90:
		fmt.Println("成绩良好")
	case score >= 60 && score < 70:
		fmt.Println("成绩及格")
	default:
		fmt.Println("不及格")
	}
}

func f3() {
	// switch后也可以直接声明/定义一个变量，分号结束，不推荐
	switch score := 90; {
	case score > 90:
		fmt.Println("成绩优秀")
	case score >= 70 && score <= 90:
		fmt.Println("成绩良好")
	case score >= 60 && score < 70:
		fmt.Println("成绩及格")
	default:
		fmt.Println("不及格")
	}
}

func f4() {
	// switch 穿透-fallthrough，如果加fallthrough则会继续执行下一个case
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	}
}

func f5() {
	// Type Switch：switch语句还可以被用于type-switch来判断某个interface变量中实际指向的变量类型
	var x interface{} // nil
	// var y = 10.0
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型: %T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是float64 型")
	}
}

// for 循环控制
// 基本语法

// 语法格式
// for 循环变量初始化；循环条件；循环变量迭代{
// 	循环操作（语句）
// }

// 执行顺序说明：
// 1.执行变量初始化，比如 i：=1
// 2.执行循环条件，比如 i <= 10
// 3.如果循环条件为真，就执行循环操作
// 4.执行循环变量迭代，比如i++
// 5.反复执行2、3、4步骤，直到循环条件为False退出for循环

// //第二种使用方式
// for 循环判断条件{
//     //循环执行语句
// }
// j ：= 1//循环变量初始化
// for j <= 10{//循环条件
//     fmt.Println("hello",j)
//     j++ //循环变量迭代
// }

// //第三种使用方式：等价for ；；{}是一个无限循环，通常配合break使用
// for{
//     //循环执行语句
// }
// k ：= 1
// for{
//     if k<= 10{
//         fmt.Println("ok"，k)
//     }else{
//         break //跳出for循环
//     }
//     k++
// }

func f6() {
	// 第一种
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 第二种
	j := 0
	for j <= 10 {
		fmt.Println("hello,j", j)
		j++ //循环变量迭代
	}

	// 第三种
	k := 0
	for {
		if k < 10 {
			fmt.Println("ok", k)
		} else {
			break //跳出循环
		}
		k++
	}
}

func f7(){
	//Golang 提供 for-range的方式，可以方便遍历字符串和数组
	//传统方式 
	var str string = "hello world"
	for i:=0; i<len(str); i++{
		fmt.Printf("%c \n",str[i])
	}

	//传统方式弊端如果str 中含有中文，会出现乱码，因为一个汉字占三个字节，需要用到切片
	//我这里还是乱码 不知道该咋解决 
	var str2 string = "hello 小王"
	str3 := []rune(str2) // 把str转成[]rune
	for i:=0; i<len(str3); i++{
		fmt.Printf("%c \n", str2[i])
	}

	// for-range 方式：是按字符方式遍历，如果有中文也是ok
	str4 := "hello 小李"
	for index, val := range str4{
		fmt.Printf("index=%d, val=%c \n",index, val)
	}
}

func f8(){
	// Go语言没有while和do..while语法。可以通过for循环来实现其使用效果

	// 1.while循环的实现
	// 循环变量初始化
	//for{
    // if循环条件表达式{
    // 	  break//跳出for循环
    // }
    // 循环操作语句
    // 循环变量迭代
	// }
	var i int
	for{
		if(i>10){
			break
		}
		fmt.Println("i的值为:",i)
		i++
	}

	//2.do..while的实现
	// 循环变量初始化
	// for{
	//     循环操作语句
	//     循环变量迭代
	//     if循环条件表达式{
	//         break//跳出for循环
	//     }
	// }

	k := 0
	for{
		fmt.Println("k的值为:",k)
		k++
		if(k>10){
			break
		}
	}

}

func f9(){
	// 跳转控制语句 break
	// break语句用于终止某个语句块的执行，用于中断当前for循环或跳出switch语句

	for i := 0;i<4;i++{
		for j := 0; j<10;j++{
			if j == 5{
				break //break 默认跳出最近的for循环
			}
			fmt.Printf("i=%d, j=%d\n",i,j)
		}
	}

	println("--------------")

	// 跳转控制语句-continue
	// continue语言用于结束本次循环，继续执行下一次循环；
	// 出现在多层嵌套的循环语句体中时，可以通过标签指明要跳过的是哪一层循环。和break规则一样
	
	for i := 0;i<4;i++{
		for j := 0; j<10;j++{
			if j %2 == 0{
				continue 
			}
			fmt.Printf("i=%d, j=%d\n",i,j)
		}
	}
}

func f10(){
	var n int = 30
	fmt.Println("ok1")
	if n > 20{
		goto label1
	}
	fmt.Println("ok2")
	label1:
	fmt.Println("ok3")
}

func f11(){
	//如果return在普通函数，则表示跳出该函数，不再执行return后面代码
	//如果return是在main函数，终止程序
	fmt.Println("ok1")
	return
	//fmt.Println("ok2")
}
func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	//f6()
	//f7()
	//f8()
	//f9()
	f11()
}

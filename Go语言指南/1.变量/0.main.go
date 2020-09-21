package main
import "fmt"
import "unsafe"
import "strconv"

func f0(){
	// 第一种使用方式
	var i int
	// 给i赋值
	fmt.Println("i=",i)
	i = 10
	fmt.Println("i=",i)

	// 第二种方式
	var num = 10.11
	fmt.Println("num=",num)

	// 第三种方式
	// 等价于 var name string name = “tom”
	// := :不能省略 否则错误
	name := "tom"
	fmt.Println("name=", name)

	// 多变量声明
	// 方式1
	// var n1, n2, n3 int
	// 方式2
	// var n1, name, n3 = 100, "tom", 888
	// 方式3
	// n1, name3, n3 := 100, "tom", 888

	// 定义全局变量
	var(
		n3 = 300
		n4 = 900
		name2 = "jack"
	)


	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(name2)
}

func f1(){
	// 数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=",i)
	// i = 1.2 //错误，不能改变数据类型

	// 变量在同一个作用域（一个函数或者在代码块）内不能重名
	// var i int = 59
	// i := 99 //错误
	// 变量 = 变量名+值+数据类型

	//Golang的变量如果没有赋初值，编译器会使用默认值，比如int默认值0 string 默认值空串 小数默认为0
}

func f2(){
	
	var i = 1
	var j = 2
	var r = i + j 

	fmt.Println("r:",r)

	var str1 = "hello"
	var str2 = "world"
	var res = str1+","+str2

	fmt.Println("res:"+res)
}

func f3(){
	// 整数类型
	var n2 int64 = 10
	//unsafe.Sizeof(n1)是unsafe包的一个函数，可以返回变量占用的字节数
	// 下面这个地方有疑问  成功解决
	fmt.Printf("n2 的类型 %T n2占用的字节数是%d",n2,unsafe.Sizeof(n2))
	//在确保程序正确运行下，尽量使用占用空间小的数据类型
	//bit:计算机中的最小存储单位。byte：计算机中基本存储单元1byte=8 bit
	//var age byte = 90

}

func f4(){
	//小数类型/浮点型
	var price float32 = 89.12
	fmt.Println("price = ",price)

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=",num3,"num4=",num4)
	//float64的精度比float32要精确
	//如果我们要保存一个精度高的数，则应该选用float64

	//Golang的浮点型默认声明位float64类型
	var num5 = 1.1
	fmt.Printf("num5的数据类型是 %T \n", num5)

	num := 5.12 //十进制数形式
	num = 5.1234e2 //10的二次方
	fmt.Println(num)
}

func f5(){
	// Golang中没有专门的字符类型，如果要存储单个字符（字母），一般用byte类保存。
	var c1 byte = 'a'
	var c2 byte = '0'
	//当我们直接输出byte值，就是输出了的对应字符的码值
	fmt.Println("c1=",c1,"c2=",c2)

	//输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2)

	var c3 int = '难'
	fmt.Printf("c3=%c c3对应码值=%d\n", c3,c3)

	var c4 byte = '\n'
	fmt.Println("c4=",c4)

	var c5 int = 22269 
	fmt.Printf("c5=%c\n",c5)

	//字符类型是可以进行运算的，相当于一个整数，运输时是按照码值运行
	var n1 = 10 +'a' // 10+97=107
	fmt.Println("n1=",n1)

	//Go语言的编码都统一成了utf-8
}

func f6(){
	//布尔类型
	var b = false
	fmt.Println("b=",b)
}

func f7(){
	//string类型
	var address string = "北京 东城区 朝阳街道"
	fmt.Println("address=",address)

	// 字符串一旦赋值了，字符串就不能修改了：在Go中字符串是不可变的
	//var str = "hello"
	//str[0] = 'a' 

	// 字符串的两种表示形式
	str2 := "abc" //(1)双引号,会识别转义字符
	fmt.Println("str2=",str2)

	//(2)反引号,以字符串的原生形式输出，包括换行和特殊字符，实现防止攻击
	str3 := `//不是单引号
		package main
		import {

		}
		func main(){

		}
	` 
	fmt.Println(str3)

	//字符串拼接方式
	var str4 = "hello"+"world"
	str4 += "haha!"
	fmt.Println(str4)
	
	//当一行字符串太长时，需要使用多行字符串，可以如下处理
	// +保留在上一行
	str5 := "hello"+"world"+
	"world"
	fmt.Println(str5)
}

func f8(){
	//基本数据类型的默认值
	var a int //0
	var b float32 //0
	var c float64 //0
	var isMarried bool //flase
	var name string //""
	fmt.Printf("a=%d,b=%v,c=%v,isMarried=%v name=%v",a,b,c,isMarried,name)
}

func f9(){
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n",i,n1,n2,n3)

	//Go中，数据类型的转换可以从 表示范围小-》表示范围大，也可以范围大-》范围小

	//被转换的是变量存储的数据（值），变量本身的数据类型并没有变化

	//在转换中，比如将int64 转换 int8【-128-127】，编译时不会报错，只是转换结果是按溢出处理，和希望的结果不一样。转换时，需要考虑范围。
	var num1 int64 = 9999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=",num2) //num2=127
}

func f10(){
	//基本数据类型和string的转换

	//基本数据类型转string类型
	//方式一：fmt.Sprintf("%参数",表达式)
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%c",myChar)
	fmt.Printf("str type %T str=%q\n",str,str)
}

func f11(){
	//方式二：使用strconv包的函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true
	var str string

	str = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type %T str=%q\n",str,str)

	//f 格式 10：表示小数位保留10位 64：表示这个小数是float64
	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n",str,str)

	//strconv 包中有一个函数Itoa
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n",str,str)

}

func f12(){
	//string类型转基本数据类型
	//使用strconv包的函数
	
	//func ParseBool(str string)(value bool,err error)
	//func ParseFloat(s string,bitSize int)(f float64,err error)
	//func ParseInt(s string,base int,bitSize int)(i int64,err error)
	//func ParseUint(s string,b int,bitSize int)(n uint64,err error)

	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v \n",b,b)

	var str2 string = "12345678"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2,10,64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n",n1,n1)
	fmt.Printf("n2 type %T n1=%v\n",n2,n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3,64)
	fmt.Printf("f1 type %T f1=%v",f1,f1)

}

func f13(){
	//指针类型
	
	//指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
	//获取变量的地址，用&
	//获取指针类型所指向的值，使用:*

	var i int = 10
	fmt.Println("i的地址=",&i)

	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Printf("ptr指向的值=%v\n",*ptr)

	var num int = 9
	var ptr1 *int = &num
	*ptr1 = 10
	fmt.Println("num=",num)//通过指针修改num值

	//值类型，都有对应的指针类型，形式位数据类型，比如int对应int，float32对应**float32 依次类推
	//值类型包括:基本数据类型int系列，float系列，bool，string、数组和结构体struct

}

func f14(){
	// 值类型：基本数据类型int系列、float系列、bool、string、数组和结构体struct
	// 引用类型：指针、slice切片、map、管道chan、interface等

	// 值类型：变量直接存储值，内存通常在栈中分配
	// 引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就变成为一个垃圾，由GC来回收。


}

func f15()  {
	// 命名规范
	// 由26个英文字母大小写，0-9，_组成
	// 数字不可以开头。
	// Golang中严格区分大小写
	// 标识符不能包含空格
	// 下划线"_"本身在Go中是一个特殊的标识符，称为空标识符。可以代表任何其他的标识符，但是它对应的值会被忽略。所以仅能被作为占位符使用，不能作为标识符使用。
	// 不能以系统保留关键字作为标识符(一共25个)，比如break，if等

	// 包名：保持package的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，不要和标准库冲突
	// 变量名、函数名、常量名：采用驼峰法
	// 如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用(注：可以简单理解成，首字母大写是公开的，首字母小写是私有的)，在golang没有public、private等关键字。
}

func main(){
	//f0()
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	//f6()
	//f7()
	//f8()
	//f9()
	//f10()
	//f11()
	//f12()
	f13()
}
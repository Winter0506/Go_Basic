package main

import "fmt"

// 接口使用注意事项和细节
// 1.接口本身不能创建实例，但是可用指向一个实现了该接口的自定义类型的变量

// Ainterface ...
type Ainterface interface {
	Say()
}

// Stu ...
type Stu struct {
	Name string
}

// Say ...
func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

// 2.接口中所有方法没有方法体
// 3.一个自定义类型将某个接口中的所有方法实现，可称为这个自定义类型实现了该接口
// 4.一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型
// 5.只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

// 6.一个自定义类型可以实现多个接口

// Binterface ...
type Binterface interface {
	Say()
}

// Cinterface ...
type Cinterface interface {
	Hello()
}

// Monster ...
type Monster struct {
}

// Hello ...
func (m Monster) Hello() {
	fmt.Println("Monster hello()")
}

// Say ...
func (m Monster) Say() {
	fmt.Println("Monster Say")
}

// 7.Golang 接口中不能有任何变量
// type AInterface interface{
//     Name string //错误
// }

/*
8.一个接口（比如A接口）可以继承多个别的接口（比如BC），如果要实现A，也必须将B，C接口的方法也全部实现
type BInterface interface{
    test01()
}
type CInterface interface{
    test02()
}
type Ainterface interface{
    BInterface
    CInterface
    test03()
}
*/

/*
9.interface类型默认是一个指针（引用类型），如果没有对interface初始化就使用，会输出nil
10.空接口没有任何方法，所以所有类型都实现了空接口，我们可以把任一变量赋值给空接口
type T interface{

}
var t T = stu
*/

func main() {
	fmt.Println("haha")

	fmt.Println("-------------")

	var stu Stu
	var a Ainterface = stu
	a.Say()

	fmt.Println("-------------")

	var i integer = 10

	var b Ainterface = i
	b.Say()

	fmt.Println("--------------")

	var monster Monster
	var b2 Binterface = monster
	var c2 Cinterface = monster
	b2.Say()
	c2.Hello()

	fmt.Println("---------------")
}

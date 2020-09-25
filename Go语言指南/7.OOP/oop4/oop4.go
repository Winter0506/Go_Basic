package main

import "fmt"

// 接口 interface
// 接口快速入门

/*
基本语法：
type 接口名 interface{
    method1(参数列表)返回值列表
    method2(参数列表)返回值列表
    ...
}
//实现接口所有方法
func(t 自定义类型) method1(参数列表)返回值列表{
    //方法实现
}
func(t 自定义类型) method2(参数列表)返回值列表{
    //方法实现
}
小结：
1.接口里的所有方法都没有方法体。体现程序设计多态和高内聚低耦合的思想
2.不需要显示的实现。没有implement关键字，只需要一个变量，含有接口中所有方法，那么这个变量就实现了这个接口
*/

// 实现多态

// USB haha
type USB interface {
	Start()
	Stop()
}

// Phone haha
type Phone struct {
	name string
}

// Start 实现usb接口方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

// Stop 实现usb接口方法
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

// Camera haha
type Camera struct {
	name string
}

// Start 实现usb接口方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

// Stop 实现usb接口方法
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

// Computer 结构体
type Computer struct {
}

// Working 方法,接收一个USB接口类型变量
func (co Computer) Working(usb USB) {
	usb.Start()
	usb.Stop()
}

func main() {
	// Golang多态
	// 变量（实例）具有多种形态，面向对象的三大特征，在Go语言中，多态是通过接口实现的。
	// 可以按照统一的接口来调用不同的实现，这时接口变量就呈现不同的形态
	fmt.Println("haha")

	fmt.Println("----------------")

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

	fmt.Println("----------------")

	var usbArr [3]USB

	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"索尼"}
	// [{vivo} {xiaomi} {索尼}]

	fmt.Println(usbArr)
}

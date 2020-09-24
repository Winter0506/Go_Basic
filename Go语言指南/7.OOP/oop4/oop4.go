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

// USB haha
type USB interface {
	Start()
	Stop()
}

// Phone haha
type Phone struct {
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
	fmt.Println("haha")

	fmt.Println("----------------")

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

}

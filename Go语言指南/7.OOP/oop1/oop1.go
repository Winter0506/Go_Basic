// VSCode的常用的快捷键
// 删除当前行crtl+shift+k
// 向上/向下复制当前行 shift+alt+↓/↑
// 补全代码 alt + /
// 添加注释和取消注释 ctrl + /
// 快速修复 alt + /
// 快速格式化代码 shift + alt + f
package main

import "fmt"

// GoLang 继承
// 嵌套匿名结构体的基本语法

// Goods 哈哈
type Goods struct {
	Name  string
	Price int
}

// Book 哈哈
type Book struct {
	Goods  // 这里就是嵌套匿名结构体Goods
	Writer string
}

// Student 学生
type Student struct {
	Name  string
	Age   int
	Score int
}

// Pupil 小学生
type Pupil struct {
	Student
}

func main() {
	fmt.Println("haha")
	fmt.Println("-----------------")
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	fmt.Println(pupil.Student.Name)
	fmt.Println(pupil.Student.Age)
	fmt.Println("-----------------")
	/*
		tom
		8
	*/

}

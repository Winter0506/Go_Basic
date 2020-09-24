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

/*
结构体可以使用嵌套匿名结构体所有的字段和方法

匿名结构体字段访问可以简化
*/

// AB ...
type AB struct {
	Name string
	age  int
}

// SayOk 说再见
func (a *AB) SayOk() {
	fmt.Println("AB SayOk",a.Name,a.age)
}

// BA extends from AB
type BA struct{
	AB
}


/*
当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，
如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分

如果一个struct嵌套一个有名的结构体，这种模式就是组合，这种方式必须带上结构体名字
*/

// D 结构体
type D struct{
	ab AB // 有名结构体 组合关系
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
	var ba BA
	ba.AB.Name = "tom"
	ba.AB.age = 18
	ba.AB.SayOk()
	// 可以简写成
	ba.Name = "tom"
	ba.age = 19
	ba.SayOk()
	fmt.Println("-----------------")

	var d D 
	d.ab.Name = "jack"
	d.ab.age = 14
	d.ab.SayOk()


}

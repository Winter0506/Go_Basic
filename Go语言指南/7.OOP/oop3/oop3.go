package main

import "fmt"

// Golang 多重继承

// 如一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承

// A1 haha
type A1 struct{
	Name string
	Age int
}

// A2 haha
type A2 struct{
	Price float64
	ID int
}

// B haha
type B struct{
	A1
	A2
}

//
func (b *B) double() {
	fmt.Println("多重继承:",b.Name,b.Age,b.Price,b.ID)
}

func main(){
	fmt.Println("haha")
	fmt.Println("----------------")
	var b B
	b.A1.Name = "牙牙乐"
	b.A1.Age = 3
	b.A2.Price = 12.44
	b.A2.ID = 1
	b.double()
}
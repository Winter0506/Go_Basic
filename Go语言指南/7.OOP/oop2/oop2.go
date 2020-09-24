package main

import "fmt"

// 嵌套匿名结构体后，也可以在创建结构体变量时，直接指定各个匿名结构体字段值


// Goods ...
type Goods struct{
	Name string
	Price float64
}

// Brand ...
type Brand struct{
	Name string
	Address string
}

// TV ...
type TV struct{
	Goods
	Brand
}

// TV2 ...
type TV2 struct{
	*Goods
	*Brand
}


func main(){
	fmt.Println("haha")

	fmt.Println("---------")

	tv := TV{Goods{"电视机001",5000.99},Brand{"海尔","山东"}}
	fmt.Println(tv.Goods.Name)

	fmt.Println("---------")

	tv2 := TV{
		Goods{
			Name:"电视机002",
			Price:5000.99,
		},
		Brand{
			Name:"海尔",
			Address:"山东",
		},
	}
	fmt.Println(tv2.Goods.Name)

	fmt.Println("---------")

	tv3 := TV2{&Goods{"电视机003",5000.99},&Brand{"海尔","山东"},}
	fmt.Println(tv3.Goods.Name)

	fmt.Println("---------")

	tv4 := TV2{
		&Goods{
			Name:"电视机004",
			Price:5000.99,
		},
		&Brand{
			Name:"海尔",
			Address:"山东",
		},
	}
	fmt.Println(tv4.Goods.Name)

	fmt.Println("---------")
}
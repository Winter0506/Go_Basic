package main

import (
	"fmt"
	"sort"
	"math/rand"
)

//接口编程的最佳实践
//sort包提供了排序切片和用户自定义数据集的函数 。
//实现对hero结构体切片的排序：sort.Sort(data Interface)
//1.声明hero结构体

// Hero ...
type Hero struct{
	Name string
	Age int
}

//2.聲明一個hero结构体切片类型
type HeroSlice []Hero

//3.实现Interface接口Len Less Swap
func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i,j int) bool{
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i,j int){
	hs[i],hs[j] = hs[j],hs[i]
}

func main(){

	fmt.Println("hello")

	fmt.Println("---------------")

	var hs HeroSlice
	for i:=0;i<10;i++{
		hero:=Hero{
			Name:fmt.Sprintf("HEro:%d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		hs = append(hs,hero)
	}

	for _,v := range hs{
		fmt.Println(v)
	}

	fmt.Println("---排序后---")
	//调用sort.Sort
	sort.Sort(hs)

	for _, v := range hs{
		fmt.Println(v)
	}

}
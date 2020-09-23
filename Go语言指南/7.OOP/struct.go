// 工厂模式实现跨包创建结构体实例
/*
如果model包的结构体变量首字母大写，引入后，直接使用

如果小写，可以用工厂模式解决
*/
// 这个不会

// package main

// type student struct{
// 	Name string
// 	score float64 // 小写
// }

// // 通过工厂模式

// //NewStudent haha
// func NewStudent(n string, s float64) *student{
// 	return &student{
// 		Name : n,
// 		score : s,
// 	}
// }

// func(s *student) GetScore() float64{
// 	return s.score
// }
package main

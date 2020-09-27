// Go语言中自带有一个轻量级的测试框架testing和自带的go test 命令来实现单元测试和性能测试
/*
案例要求：
1.编写一个Monster结构体，字段Name、Age、Skill
2.给Monster绑定方法Store，可以将一个Monster变量(对象),序列化后保存到文件中
3.给Monster绑定方法ReStore，可以将一个序列化的Monster从文件中读取，并反序列化为Monster对象，检查反序列化，名字正确
4.编程测试用例文件store_test.go ,编写测试用例函数TestStore和TestRestore进行测试
*/


// import(
// 	"encoding/json"
// 	"io/ioutil"
// 	"fmt"
// )

// Monster ...
// type Monster struct {
// 	Name  string
// 	Age   int
// 	Skill string
// }

// Store ...
// func (*Monster) Store() bool{
// 	data,err := json.Marshal(&Monster) // 不对
// 	if err != nil{
// 		fmt.Println("marshal err=",err)
// 		return false
// 	}
// 	filePath := "Go_Basic/Go语言指南/9.单元测试/monster.ser"
// 	err = ioutil.WriteFile(filePath,data,0666)
// 	if err != nil{
// 		fmt.Println("write file err=",err)
// 		return false
// 	}
// 	return true
// }

// func (this *Monster) ReStore() bool{
// 	filePath := "Go_Basic/Go语言指南/9.单元测试/monster.ser"
// 	data,err := ioutil.ReadFile(filePath)
// 	if err != nil{
// 		fmt.Println("Reader file err=",err)
// 		return false
// 	}

// 	err = json.Unmarshal(data,this)
// 	if err != nil{
// 		fmt.Println("Unmarshal err=",err)
// 		return false
// 	}
// 	return true
// }

package main

import "fmt"

func main(){
	fmt.Println("-------")
}


// import(
// 	"testing"
// )

// func TestStore(t *testing.T){
// 	monster := &Monster{
// 		Name : "红孩儿",
// 		Age : 10,
// 		Skill : "吐火",
// 	}

// 	res := monster.Store()
// 	if !res{
// 		t.Fatalf("monster.Store() 错误,希望为=%v 实际为=%v",true,res)
// 	}
// 	t.Logf("monster.Store() 测试成功!")
// }

// func TestReStore(t *testing.T){
// 	var monster = &Monster{}
// 	res := monster.ReStore()
// 	if !res {
// 		t.Fatalf("monster.ReStore() 错误,希望为=%v 实际为=%v","红孩儿",monster.Name)
// 	}
// 	t.Logf("monster.ReStore() 测试成功!")
// }


/*
//执行命令：go test -v monster_test.go monster.go
=== RUN   TestStore
    TestStore: monster_test.go:18: monster.Store() 测试成功！
--- PASS: TestStore (0.00s)                                                                                  tom"}]
=== RUN   TestReStore
    TestReStore: monster_test.go:31: monster.ReStore() 测试成功！
--- PASS: TestReStore (0.00s)
PASS
ok      command-line-arguments  0.957s
*/
package main

// 命令行参数
import (
	"flag"
	"fmt"
	"os"
)

// os.Args是一个string的切片，用来存储所有的命令行参数
func main() {
	fmt.Println("命令行所有参数:", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
	/*
		命令行所有参数: 1
		args[0]=C:\Users\15697\AppData\Local\Temp\go-build265279530\b001\exe\4.args.exe
	*/

	// flag包用来解析命令行参数
	var user string
	var pwd string
	var host string
	var port string

	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.StringVar(&port, "p", "3306", "端口号,默认为3306")
	flag.Parse() // 此方法必须调用

	fmt.Printf("user=%v pwd=%v host=%v port=%v",user,pwd,host,port)
	//测试 go run XXX.go -u root -pwd 123456 -h 127.0.0.1 -p 3366

	/*
	args[0]=C:\Users\15697\AppData\Local\Temp\go-build238994914\b001\exe\4.args.exe
	args[1]=-u
	args[2]=root
	args[3]=-pwd
	args[4]=123456
	args[5]=-h
	args[6]=127.0.0.1
	args[7]=-p
	args[8]=3366
	user=root pwd=123456 host=127.0.0.1 port=3366
	*/
}

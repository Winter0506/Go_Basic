// 写文件操作应用实例

// 基本介绍os.OpenFile函数
// 创建一个新文件,写入内容5句"hello world"

package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	filePath := "Go_Basic/Go语言指南/8.文件操作/write/write.txt"
	file,err := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE, 0666)

	if err != nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}

	defer file.Close()

	str := "hello world\n"
	writer := bufio.NewWriter(file)
	for i:=0; i<5; i++{
		writer.WriteString(str)
	}
	// writer是带缓存，调用writerString方法时，内容先写入到缓存的，
	// 所有需要调用Flush方法，将缓冲的数据真正写入到文件中，否则文件中会没有数据
	writer.Flush()
	
}


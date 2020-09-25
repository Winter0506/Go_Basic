package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)

func main(){
	file, err := os.Open("/home/wangqichao/wangqichaowork/GoProject/Go_Basic/Go语言指南/8.文件操作/example.txt")
	if err != nil{
		fmt.Println("open file err=",err)
	}

	defer file.Close()

	const(
		defaultBufsize = 4096
	)

	reader := bufio.NewReader(file)

	for{
		str,err := reader.ReadString('\n')
		if err == io.EOF{ // io.EOF 表示文件的末尾
			break
		}
		fmt.Println(str)
	}

	fmt.Println("文件读取结束")

	fmt.Println("-------------")

	fmt.Printf("file=%v\n",file) // file=&{0xc000052180}

	fmt.Println("-------------")


	err = file.Close()
	if err != nil{
		fmt.Println("close file err=",err)
	}

	// 读取文件的内容并显示在终端（使用ioutil一次将整个文件读入到内存中），
	// 这种方式适用于文件不大的情况，相关方法和函数（）ioutil.ReadFile

	file1 := "Go_Basic/Go语言指南/8.文件操作/example.txt"
	content, err := ioutil.ReadFile(file1)
	if err != nil{
		fmt.Printf("read file err=%v",err)
	}
	fmt.Printf("%v",string(content))

	fmt.Println("-------------")

	

}
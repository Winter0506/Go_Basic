// 写文件操作应用实例

// 基本介绍os.OpenFile函数
// 创建一个新文件,写入内容5句"hello world"

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

// 判断文件是否存在

// PathExists ...
func PathExists(path string) (bool,error){
	_, err := os.Stat(path)
	if err == nil{
		return true, nil
	}
	if os.IsNotExist(err){
		return false, nil
	}
	return false,err
}

func main() {
	filePath := "Go_Basic/Go语言指南/8.文件操作/write/write.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "hello world\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// writer是带缓存，调用writerString方法时，内容先写入到缓存的，
	// 所有需要调用Flush方法，将缓冲的数据真正写入到文件中，否则文件中会没有数据
	writer.Flush()

	// 在原来的内容追加内容 hello everyone
	str1 := "hello everyone\n"
	writer1 := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer1.WriteString(str1)
	}
	writer1.Flush()

	// 打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句 “hello everything”

	file1,err1 := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)


	if err1 != nil{
		fmt.Printf("open file err=%v\n",err1)
		return
	}

	defer file1.Close()
	// 读取文件内容
	reader := bufio.NewReader(file1)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str2 := "hello everything\n"
	writer2 := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer2.WriteString(str2)
	}

	writer2.Flush()

	// 编写一个程序，将一个文件的内容，写入到另外一个文件。这两个文件已经存在。
	// 使用ioutil.ReadFile/ioutil.WriteFile 完成写文件的任务

	filePath1 := "Go_Basic/Go语言指南/8.文件操作/write/write1.txt"

	data, err := ioutil.ReadFile(filePath)

	if err != nil{
		fmt.Printf("open file err=%v\n",err)
	}

	err = ioutil.WriteFile(filePath1,data,0666)

	if err != nil{
		fmt.Printf("write file error = %v\n",err)
	}

	// 判断文件是否存在
	fmt.Println("--------------")
	var exist bool
	exist, _ = PathExists(filePath)
	if exist{
		fmt.Println("存在文件")
	}


}

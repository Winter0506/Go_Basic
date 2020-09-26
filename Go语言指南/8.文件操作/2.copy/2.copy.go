package main

import(
	"fmt"
	"os"
	"io"
	"bufio"
)

// CopyFile ...
func CopyFile(dstFileName string, srcFileName string)(written int64,err error){
	srcFile, err := os.Open(srcFileName)
	if err != nil{
		fmt.Printf("open file err=%v\n",err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE,0666)

	if err != nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer,reader)
}

func main(){
	srcFile := "Go_Basic/Go语言指南/8.文件操作/2.copy/test.txt"
	dstFile := "Go_Basic/Go语言指南/8.文件操作/2.copy/test1.txt"

	_,err := CopyFile(dstFile,srcFile)
	if err == nil{
		fmt.Printf("拷贝完成\n")
	}else{
		fmt.Printf("拷贝错误 err=%v\n",err)
	}
}
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CharCount ...
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	fileName := "Go_Basic/Go语言指南/8.文件操作/2.copy/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 读到文件末尾
			break
		}
		// 为兼容中文字符,将str转成[]rune
		str1 := []rune(str)
		for _, v := range str1 {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格个数为=%v 其他字符个数为=%v\n", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}

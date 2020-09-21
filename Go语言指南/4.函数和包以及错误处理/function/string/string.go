package main

import (
	"strings";
	"fmt";
	"strconv"
)

func main() {
	// 1.统计字符串的长度，按照字节len(str)
	str := "hello北京"
	fmt.Println("str len = ", len(str)) // 1

	fmt.Println("----------")

	// 2.字符串遍历,同时处理有中文的问题 r:=[]rune(str)
	str2 := "hello上海"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	fmt.Println("----------")

	// 3.字符串转整数:   这儿有点儿乱哈！！！！
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
		fmt.Println("转成的结果是", n)
	} else {
		fmt.Println("转成的结果是", n)
	}
	n1, err1 := strconv.Atoi("100")
	if err != nil {
		fmt.Println("转换错误", err1)
		fmt.Println("转成的结果是", n1)
	} else {
		fmt.Println("转成的结果是", n1)
	}
	fmt.Println("----------")
	/*
		转换错误 strconv.Atoi: parsing "hello": invalid syntax
		转成的结果是 0
		转换错误 <nil>
		转成的结果是 100
	*/

	// 4.整数转字符串
	str3 := strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n", str3, str3)
	fmt.Println("-----------")

	// 5.字符串转 []byte  var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n",bytes)
	fmt.Println("------------")

	// 6.[]byte转字符串： str = string([]byte{97，98，99})
	str = string([]byte{97,98,99})
	fmt.Printf("str=%v\n",str)
	fmt.Println("------------")

	// 7.10进制转2，8，16进制： str = strconv.FormatInt(123,2) //2->8,16
	str = strconv.FormatInt(16,2)
	fmt.Printf("16对应的二进制是=%v\n",str)
	str = strconv.FormatInt(16,16)
	fmt.Printf("16对应的十六进制是=%v\n",str)
	fmt.Println("------------")

	// 8.查找子串是否在指定的字符串中：strings.Contains("seafood","foo")
	b := strings.Contains("seafood","mary")
	fmt.Printf("b=%v\n",b)
	b = strings.Contains("seafood","foo")
	fmt.Printf("b=%v\n",b)
	fmt.Println("------------")

	// 9.统计一个字符串有几个指定的子串：strings.Count("ceheese","e") //4
	num := strings.Count("ceheese","e")
	fmt.Printf("num=%v\n",num)
	fmt.Println("------------")

	// 10.不区分大小写的字符串比较（==是区分字母大小写的）
	b = strings.EqualFold("abc","Abc")
	fmt.Printf("b=%v",b) //true
	fmt.Println("结果","abc" == "Abc") // false 区分大小写
	fmt.Println("-------------")

	// 11.返回子串在字符串第一次出现的index值，如果没有返回-1 ：
	index := strings.Index("NLT_abcabcabc","abc")
	fmt.Println(index)
	fmt.Println("-------------")

	// 12.返回字串在字符串最后一次出现的index，如果没有返回-1
	index = strings.LastIndex("gogocomego","go") //8
	fmt.Printf("index=%v\n",index)
	fmt.Println("-------------")

	// 13.将指定的字串替换成另外一个字串
	str4 := "java java hello java"
	str5 := strings.Replace(str4, "java", "go", -1) //-1表示全部替换，也可以指定希望替换几个
	fmt.Printf("str4=%v,str5=%v\n",str4,str5)
	fmt.Println("-------------")

	// 14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,word,ok",",")
	for i:=0;i<len(strArr);i++{
		fmt.Printf("str[%d]=%v\n",i,strArr[i])
	}
	fmt.Printf("strArr=%v\n",strArr)
	fmt.Println("-------------")

	// 15.将字符串的字母进行大小写的转换
	str6 := "Hello World"
	str6 = strings.ToLower(str6)
	fmt.Printf("str=%v\n",str6)
	str6 = strings.ToUpper(str6)
	fmt.Printf("str=%v\n",str6)
	fmt.Println("-------------")

	// 16.将字符串左右两边的空格去掉
	str7 := strings.TrimSpace(" this is a space  ")
	fmt.Printf("str7=%q\n",str7)
	fmt.Printf("str7=%v\n",str7)
	// str7="this is a space"
	// str7=this is a space
	fmt.Println("-------------")

	// 17.将字符串左右两边指定的字符去掉
	str8 := strings.Trim("go hegollogo     ","go ")
	fmt.Printf("str=%q\n",str8)   //str="hegoll"
	fmt.Println("-------------")

	// 18.将字符串左边指定的字符去掉
	str9 := strings.TrimLeft("go hegollogo     ","go ")
	fmt.Printf("str=%q\n",str9)   //str="hegollogo     "
	fmt.Println("-------------")

	// 19.将字符串右边指定的字符去掉
	str10 := strings.TrimRight("go hegollogo     ","go ")
	fmt.Printf("str=%q\n",str10)   //str="go hegoll"    "
	fmt.Println("-------------")

	// 20.判断字符串是否以指定的字符串开头
	b = strings.HasPrefix("ftp://192.168.0.1","ftp")
	fmt.Println(b)
	fmt.Println("-------------")

	// 21.判断字符串是否以指定的字符串结束
	b = strings.HasSuffix("hello.txt","txt")
	fmt.Println(b)
	fmt.Println("-------------")

}

package add
import "fmt"

// 在给一个文件打包时，该包对应一个文件夹，包名和文件夹名一致，一般为小写字母

func Sum(n1 float64, n2 float64) float64{
	var res float64
	res = n1 + n2
	fmt.Println("n1 + n2 =",res)
	return res
} 

// 函数实例
func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误")
	}
	return res
}

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'
	result := cal(n1, n2, operator) //调用函数
	fmt.Println("result=", result)
}
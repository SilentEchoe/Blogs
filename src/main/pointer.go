package main
// 导入系统包
import (
    "flag"
    "fmt"
)
// 定义命令行参数
var mode = flag.String("mode", "", "process mode")
func main() {

	// 创建一个string 类型的指针
	str := new(string)

	*str = "Go语言教程"

	fmt.Println(*str)
    // 解析命令行参数
    flag.Parse()
    // 输出命令行参数
    fmt.Println(*mode)
}
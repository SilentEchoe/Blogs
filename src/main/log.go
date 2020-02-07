// 类型断言 用于检查接口类型变量所持有的值是否实现期望的接口 或具体的类型

//value,ok := x.(T)
package main

import (
	"fmt"
	"os"
	"errors"
)

// 声明日志写入器接口
type LogWriter interface{
	Write(data interface{}) error
}

// 日志构造函数
type Logger struct{
	writerList []LogWriter
}

// 注册一个日志写入器

func (l *Logger) RegisterWriter(writer LogWriter){
	l.writerList = append(l.writerList, writer)
}

func (l *Logger) Log(data interface{}){
	
	for _, writer := range l.writerList{
		// 将日志输出到每一个写入器中
		writer.Write(data)
	}
}

func  main()  {
	var a int
	a = 10
	// 验证类型
	getType(a)
	
}


func getType(a interface{}){
	switch a.(type){
		case int:
			fmt.Println("the type of a is int")
		case string:
			fmt.Println("the type of a is string")
		case float64:
			fmt.Println("the type of a is float64")
		default:
			fmt.Println("unknown type")
	}
}
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


type fileWriter struct{
	file *os.File
}


// 日志构造函数
type Logger struct{
	writerList []LogWriter
}

func (f *fileWriter) SetFile(filename string) (err error){
	// 如果文件已打开，关闭前一个文件
	if f.file != nil{
		f.file.Close()
	}

	// 创建一个文件并保存文件句柄
	f.file, err = os.Create(filename)

	return err

}

// 实现LogWriter的Write()方法
func (f * fileWriter) Write(data interface{}) error{
	// 日志文件可能没有创建成功
	if f.file == nil{
		// 日志文件没有准备好
		return errors.New("file not creadted")
	}

	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n",data)

	// 将数据以字节数组写入文件中

	_, err := f.file.Write([]byte(str))

	return err
}

// 创建文件写入器实例
func newFileWriter() *fileWriter{
	return &fileWriter{}
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
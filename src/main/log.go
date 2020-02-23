package main

import "log"

func init()  {
	// 设置日志项的前缀
	log.SetPrefix("TRACE：")
	// 标识用于控制日志项信息
	// Ldate 代表日期
	// Lmicroseconds 代表毫秒级别相应时间
	// Llongfile 代表完整的文件名和行号
	log.SetFlags(log.Ldate | log.Lmicroseconds|log.Llongfile)
}

func main()  {
	log.Println("message")
	log.Fatalln("fatal message")
	log.Panicln("Panic message")
}

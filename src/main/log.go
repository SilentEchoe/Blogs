package main

import "log"

func init()  {
	// 设置
	log.SetPrefix("TRACE：")
	log.SetFlags(log.Ldate | log.Lmicroseconds|log.Llongfile)
}

func main()  {
	log.Println("message")
	log.Fatalln("fatal message")
	log.Panicln("Panic message")
}

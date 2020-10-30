package main

import (
	"fmt"
	"runtime"
	"time"
)

func bar() {
	fmt.Println("bar:", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("foo:", runtime.NumGoroutine())
	go bar()
	time.Sleep(100 * time.Millisecond)
}

func main() {
	fmt.Println("main:", runtime.NumGoroutine())
	go foo()
	time.Sleep(100 * time.Millisecond)

	fmt.Println(runtime.GOROOT())
}

package main

import (
	"fmt"
	"reflect"
)

func main() {

	for {
		var chanSum = make(chan int, 4)
		for i := 1; i <= 4; i++ {
			chanSum <- i
		}
		close(chanSum)
		for v := range chanSum {
			//fmt.Println(typeof(v))
			go doWork(v)

		}

	}
}

func doWork(i int) {
	fmt.Println(i)
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

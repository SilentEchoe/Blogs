package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	arrOne := arr[0:3]
	arr[1] = 4
	fmt.Println(arrOne)
}

func doWork() {

}

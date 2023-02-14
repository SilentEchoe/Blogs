package main

import (
	"fmt"
	"time"
)

func main() {

	getTime := time.Now().AddDate(0, -1, 0).Unix()
	fmt.Println(getTime)
}

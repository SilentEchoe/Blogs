package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(hypot(1, 2))
}

// (x, y float64) 中 x y 为形参
// float64 为返回参数类型
func hypot(x, y float64) float64 {

	return math.Sqrt(x*x + y*y)
}

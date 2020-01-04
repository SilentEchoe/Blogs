package model

import (
	"fmt"
	"math"
)

var hp string = "Good"

func hypot(x, y float64, z string) float64 {
	fmt.Print(z)
	return math.Sqrt(x*x + y*y)
}

package main

import (
	"math"
)

const tolerance = 0.00001

// 判断两个浮点数是否接近
func isGoodEnough(guess, x float64) bool {
	return math.Abs(guess*guess-x) < tolerance
}

// 改进猜测
func improve(guess, x float64) float64 {
	return (guess + x/guess) / 2
}

// 递归尝试求平方根
func sqrtIter(guess, x float64) float64 {
	if isGoodEnough(guess, x) {
		return guess
	}
	return sqrtIter(improve(guess, x), x)
}

// 外部调用接口
func Sqrt(x float64) float64 {
	return sqrtIter(1.0, x) // 初始猜测为 1.0
}

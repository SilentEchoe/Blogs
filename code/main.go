package main

import (
	"fmt"
	"time"
)

func main() {
	//maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})

	start, end := GetUTCRangeFromStartOfLastMonthToNowInLocation(8)
	fmt.Printf("UTC+8 中从上月第一天到现在，对应的 UTC 范围是：[%d, %d]\n", start, end)
}

// LeetCode.11 盛最多水的容器
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	ans := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		fmt.Println(area)
		ans = max(ans, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GetUTCRangeFromStartOfLastMonthToNowInLocation 获取“指定时区”中从上月第一天零点到当前时间的 UTC 范围
// offsetHours：时区偏移（如 8 表示 UTC+8）
func GetUTCRangeFromStartOfLastMonthToNowInLocation(offsetHours int) (startUTC, endUTC int64) {
	// 构造时区
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", offsetHours), offsetHours*3600)

	// 当前时间（该时区）
	now := time.Now().In(loc)

	// 计算“上个月”的年份与月份
	year, month := now.Year(), now.Month()
	if month == time.January {
		year -= 1
		month = time.December
	} else {
		month -= 1
	}

	// 得到该时区下的“上个月第一天零点”
	startLocal := time.Date(year, month, 1, 0, 0, 0, 0, loc)

	// UTC 时间戳
	startUTC = startLocal.UTC().Unix()
	endUTC = now.UTC().Unix()

	return startUTC, endUTC
}

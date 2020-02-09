package main

import (
	"fmt"
	"sort"
)

// 将[]string定义为MyStringList类型

type MyStringList []string

// 获取元素数量
func (m MyStringList) Len() int{
	return len(m)
}

// 比较元素大小
func (m MyStringList) Less(i,j int) bool  {
	return m[i]<m[j]
}

// 实现元素交换

func (m MyStringList) Swap(i,j int)  {
	m[i],m[j] = m[j],m[i]
}

func main()  {
	//准备一个切片

	names := []string{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Strings(names)

	for _, v := range names {
		fmt.Printf("%s\n",v)
	}




}

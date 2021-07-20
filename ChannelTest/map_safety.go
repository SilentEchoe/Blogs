/*
 关于 Map 在线程中,如何达成线程安全的代码测试
	1.使用 sync.map
	2.使用互斥锁
	3.使用读写锁
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var sm sync.Map
var rw sync.RWMutex
var wg sync.WaitGroup

func main() {

	// 方案一
	//wg.Add(1)
	//go mapAdd1()
	//go mapAdd2()
	//wg.Wait()
	//sm.Range(func(k, v interface{}) bool {
	//	fmt.Print(k)
	//	fmt.Print(":")
	//	fmt.Print(v)
	//	fmt.Println()
	//	return true
	//})

	// 方案二

	// 互斥锁共享变量
	mutualMap := make(map[int]string)

	go mapWork1(mutualMap)
	go mapWork2(mutualMap)
	time.Sleep(2000)
	for k, v := range mutualMap {
		fmt.Println(k, v)
	}
}

// 方案1 使用 sync.map
func mapAdd1() {

	sm.Store(1, "a")
	sm.Store(2, "b")
	sm.Store(3, "c")

	// LoadOrStore 方法,获取或者保存
	// 参数是一对key：value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；不存在则store，返回该value 和false
	if vv, ok := sm.LoadOrStore(4, "d"); ok {
		fmt.Println(vv)
	}

	defer wg.Done()

}

func mapAdd2() {
	if vv, ok := sm.LoadOrStore(4, "e"); ok {
		fmt.Println(vv)
	}
	if vv, ok := sm.LoadOrStore(5, "f"); ok {
		fmt.Println(vv)
	}
	if vv, ok := sm.LoadOrStore(6, "g"); ok {
		fmt.Println(vv)
	}
}

// 方案2 使用互斥锁
func mapWork1(mutualMap map[int]string) {
	rw.Lock()
	defer rw.Unlock()
	mutualMap[1] = "a"
}

func mapWork2(mutualMap map[int]string) {
	rw.Lock()
	defer rw.Unlock()
	mutualMap[1] = "b"
	mutualMap[2] = "c"
}

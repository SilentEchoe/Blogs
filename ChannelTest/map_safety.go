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
)

var sm sync.Map
var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go mapAdd1()
	go mapAdd2()

	wg.Wait()
	sm.Range(func(k, v interface{}) bool {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
		return true
	})
}

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

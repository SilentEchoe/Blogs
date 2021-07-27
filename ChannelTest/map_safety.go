/*
 关于 Map 在线程中,如何达成线程安全的代码测试
	1.使用 sync.map
	2.使用互斥锁
	3.使用读写锁
	4.锁会降低性能,如果要使用锁,尽量减少锁的粒度和锁持有的时间 分片加锁 https://github.com/orcaman/concurrent-map
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
	//// 互斥锁共享变量
	//mutualMap := make(map[int]string)
	//go mapWork1(mutualMap)
	//go mapWork2(mutualMap)
	//time.Sleep(2000)
	//for k, v := range mutualMap {
	//	fmt.Println(k, v)
	//}

	// 方案三
	mutualMap := make(map[int]string)
	go mapRwWork3(mutualMap)
	go mapRwWork2(mutualMap)

	time.Sleep(2000)
	go mapRwWork1(mutualMap)
	time.Sleep(3000)
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

// 方案3 使用读写锁
func mapRwWork1(m map[int]string) {
	// 读锁

	fmt.Println("读锁")
	rw.RLock()
	defer rw.RUnlock()

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func mapRwWork2(m map[int]string) {
	fmt.Println("写锁2")
	rw.Lock()
	defer rw.Unlock()
	m[1] = "a"
	m[2] = "b"
}

func mapRwWork3(m map[int]string) {
	fmt.Println("写锁3")
	rw.Lock()
	defer rw.Unlock()
	m[3] = "c"
	m[4] = "d"
}

package main

import "sync"

func main() {
	var m = make(map[int]int, 10) // 初始化一个map
	go func() {
		for {
			m[1] = 1 //设置key
		}
	}()

	go func() {
		for {
			_ = m[2] //访问这个map
		}
	}()
	select {}
}

type RWMap struct {
	sync.RWMutex
	m map[int]int
}

func (m *RWMap) Get(k int) (int, bool) {
	// 从 map 读一个值
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *RWMap) Set(k int, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}

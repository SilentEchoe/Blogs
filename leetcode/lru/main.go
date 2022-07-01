package main

import (
	"container/list"
)

//LRU 算法可以认为是相对平衡的一种淘汰算法。LRU 认为，如果数据最近被访问过，那么将来被访问的概率也会更高。
//LRU 算法的实现非常简单，维护一个队列，如果某条记录被访问了，则移动到队尾，那么队首则是最近最少访问的数据，淘汰该条记录即可。

type Cache struct {
	maxBytes  int64                    //链表最大值
	nbytes    int64                    //当前字节大小
	ll        *list.List               //双向链表
	cache     map[string]*list.Element //键为字符串，值为双向链表中对应节点的指针
	OnEvicted func(key string, value Value)
}

// 双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value
}

// Value 记录链表的长度
type Value interface {
	Len() int
}

// 实现接口
func (c *Cache) Len() int {
	return c.ll.Len()
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get 如果键对应的链表存在,则对应节点移动到队尾，并返回查找的值
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		// 移动到最前面
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除链表最前面一个节点
func (c *Cache) RemoveOldest() {
	// 得到队首节点，从链表中删除
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		// 从map中也删除
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)

		//更改字节大小
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}

}

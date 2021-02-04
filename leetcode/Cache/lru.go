package main

//维护一个队列，如果某条记录被访问了，则移动到队尾，那么队首则是最近最少访问的数据，淘汰该条记录
import "fmt"

type Cachelist struct {
	data interface{}
	last *Cachelist
	next *Cachelist
}

func main() {
	var c = NewCachelist()
	c.SetCache(1)
	c.SetCache(2)
	c.SetCache(3)
	c.SetCache(4)
	c.SetCache(5)
	c.SetCache(6)
	c.SetCache(7)
	c.SetCache(8)
	c.MoveLastNode(3)
	c.RemovefirstNode()
	c.Traverse()
}

// new 会初始化一个指针类型的结构体，初始化的值为零值
func NewCachelist() *Cachelist {
	return new(Cachelist)

}

func (c *Cachelist) SetCache(i interface{}) {
	n := NewCachelist()
	n.data = i
	point := c
	for point.next != nil {
		point = point.next
	}
	point.next = n
	n.last = point

}

//遍历链表
func (c *Cachelist) Traverse() {
	point := c.next
	for nil != point {
		fmt.Println(point.data)
		point = point.next
	}
	fmt.Println("--------done----------")
}

//查到最后一个结点
func (c *Cachelist) lastNode() *Cachelist {
	point := c
	for point.next != nil {
		point = point.next
	}
	return point
}

// 将某个结点移动到最后
func (c *Cachelist) MoveLastNode(i interface{}) {
	point := c
	for point.next != nil {
		if point.data == i {
			break
		}
		point = point.next
	}
	point.last.next = point.next
	point.next.last = point.last

	var p = c.lastNode()
	point.next = nil
	point.last = p
	p.next = point
}

// 移除第一个node
func (c *Cachelist) RemovefirstNode() {
	if c.next != nil {
		c.next.last = nil
		c.next = c.next.next
	}
}

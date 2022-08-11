package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func main() {
	nodeFive := &listNode{
		Val:  5,
		Next: nil,
	}

	nodeFour := &listNode{
		Val:  4,
		Next: nodeFive,
	}

	nodeThree := &listNode{
		Val:  3,
		Next: nodeFour,
	}

	nodeTwo := &listNode{
		Val:  2,
		Next: nodeThree,
	}

	nodeHead := &listNode{
		Val:  1,
		Next: nodeTwo,
	}

	nodeHead.addAtHead(6)
	var index = nodeHead.get(1)
	fmt.Println(index)
	fmt.Println(index)
}

// 获取链表中的第 index 个节点的值。如果索引无效，返回 -1
func (head *listNode) get(index int) int {
	current := head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	if current == nil {
		return -1
	}

	return current.Val
}

//在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (head *listNode) addAtHead(v int) {
	var newHead = &listNode{Val: v, Next: nil}
	current := head
	for current != nil {
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	current.Next = newHead
	Reverse(current)
	Reverse(current)
	head = current

}

func Reverse(head *listNode) *listNode {
	if head.Next == nil {
		return head
	}
	last := Reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

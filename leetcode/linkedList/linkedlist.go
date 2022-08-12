package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

type head struct {
	headNode *listNode
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

	head := &head{headNode: nodeHead}

	head.addAtHead(6)

	var index = head.headNode.get(1)
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
func (head *head) addAtHead(v int) {
	var newHead = &listNode{Val: v, Next: head.headNode}
	head.headNode = newHead
}

// 将值为 val 的节点追加到链表的最后一个元素
func (head *head) addAtTail(v int) {
	curr := head.headNode
	var newHead = &listNode{Val: v, Next: nil}
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newHead
}

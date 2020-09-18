package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	nodeLink := new(ListNode)
	nodeLink.Val = 2

	node1 := new(ListNode)
	node1.Val = 4
	nodeLink.Next = node1

	node2 := new(ListNode)
	node2.Val = 3
	node1.Next = node2

	nodeLink2 := new(ListNode)
	nodeLink2.Val = 5

	node11 := new(ListNode)
	node11.Val = 6
	nodeLink2.Next = node11

	node22 := new(ListNode)
	node22.Val = 4
	node11.Next = node22

	//addTwoNumbers(nodeLink, nodeLink2)
	nowlistNode := addTwoNumbers(nodeLink, nodeLink2)
	for {
		if nowlistNode != nil {
			fmt.Println(nowlistNode.Val)
			nowlistNode = nowlistNode.Next
		} else {
			break
		}
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var strl1 = LinkedToString(l1)

	var strl2 = LinkedToString(l2)

	int1, _ := strconv.ParseInt(strl1, 10, 64)
	int2, _ := strconv.ParseInt(strl2, 10, 64)
	str3 := strconv.FormatInt(int1+int2, 10)

	return NewNode(str3)
}

func LinkedToString(l1 *ListNode) string {
	var strl1 string
	nowlistNode := l1
	for {
		if nowlistNode != nil {
			str21 := strconv.Itoa(nowlistNode.Val)
			strl1 = str21 + strl1
			nowlistNode = nowlistNode.Next
		} else {
			break
		}
	}
	return strl1
}

// 添加节点到链表尾部
func (node *ListNode) AddNode(int2 int) {
	for {
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}
	r := new(ListNode)
	r.Val = int2
	node.Next = r
}

func NewNode(str string) *ListNode {
	r := new(ListNode)
	str1 := str[len(str)-1 : len(str)]
	intr, _ := strconv.Atoi(str1)
	r.Val = intr

	for i := len(str) - 1; i > 0; i-- {
		str2 := str[i-1 : i]
		int2, _ := strconv.Atoi(str2)
		r.AddNode(int2)

	}

	return r
}

package main

import (
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

	addTwoNumbers(nodeLink, nodeLink)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var strl1 string
	var strl2 string
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

	nowlistNode2 := l2

	for {
		if nowlistNode2 != nil {
			str21 := strconv.Itoa(nowlistNode2.Val)
			strl2 = str21 + strl2
			nowlistNode2 = nowlistNode2.Next
		} else {
			break
		}
	}

	int1, _ := strconv.ParseInt(strl1, 10, 64)
	int2, _ := strconv.ParseInt(strl2, 10, 64)

	println(int1 + int2)
	return nil
}

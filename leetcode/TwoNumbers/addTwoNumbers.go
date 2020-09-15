package main

import (
	"fmt"
	"os"
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
	nowlistNode := l1

	for {
		if nowlistNode != nil {
			string := strconv.Itoa(l1.Val)
			strl1 = string + strl1

			nowlistNode = nowlistNode.Next
		} else {
			break
		}
	}

	int64, err := strconv.ParseInt(strl1, 10, 64)

	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	println(int64)
	return nil
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nodeThree := &ListNode{
		Val:  2,
		Next: nil,
	}

	nodeTwo := &ListNode{
		Val:  1,
		Next: nodeThree,
	}

	nodeHead := &ListNode{
		Val:  1,
		Next: nodeTwo,
	}
	//fmt.Println(*deleteDuplicatesTwo(nodeHead))

	middleNode(nodeHead)
}

/*
单向链表
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
*/
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		// 全部删除完再移动到下一个元素
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

// 请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
func deleteDuplicatesTwo(head *ListNode) *ListNode {
	current := head
	for current != nil {
		// 全部删除完再移动到下一个元素
		for current.Next != nil && current.Val == current.Next.Val {
			current = current.Next.Next
		}
	}
	return head
}

//输入：[1,2,3,4,5]
//输出：此列表中的结点 3 (序列化形式：[3,4,5])
//返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
//注意，我们返回了一个 ListNode 类型的对象 ans，这样：
//ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.
// 解题思路 快慢指针
func middleNode(head *ListNode) *ListNode {
	current := head
	count := 0
	for current != nil {
		count++
		// 全部删除完再移动到下一个元素
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	fmt.Println(count)
	return current
}

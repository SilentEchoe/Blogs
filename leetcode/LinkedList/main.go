package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nodeFive := &ListNode{
		Val:  5,
		Next: nil,
	}

	nodeFour := &ListNode{
		Val:  4,
		Next: nodeFive,
	}

	nodeThree := &ListNode{
		Val:  3,
		Next: nodeFour,
	}

	nodeTwo := &ListNode{
		Val:  2,
		Next: nodeThree,
	}

	nodeHead := &ListNode{
		Val:  1,
		Next: nodeTwo,
	}
	//fmt.Println(*deleteDuplicatesTwo(nodeHead))

	middleNode(nodeHead)

	removeNthFromEnd(nodeHead, 2)
}

/* 合并两个有序链表
 */

func mergeTwoLists(node1 *ListNode, node2 *ListNode) *ListNode {
	// 虚拟头节点
	var dummy = &ListNode{}
	p := dummy
	var p1 = node1
	var p2 = node2

	for p1 != nil && p2 != nil {
		// 比较 p1 和 p2 两个指针
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}

		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}

// 合并 K 个有序链表
func mergeKLists(nodes []ListNode) *ListNode {
	if len(nodes) == 0 {
		return nil
	}

	// 虚拟头结点
	var dummy = &ListNode{}
	_ = dummy
	// 使用 优先级队列（二叉堆） 把链表结点放入一个最小堆

	return nil
}

// 单链表的倒数第K个结点
// 可以用快慢指针
func findFormEnd(head *ListNode, k int) *ListNode {
	var p1 = head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	var p2 = head
	for p1 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	return p2
}

func removeNthFormEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head
	var x = findFormEnd(dummy, n+1)
	x.Next = x.Next.Next
	return dummy.Next
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
//fast/slow 刚开始均指向链表头节点，然后每次快节点走两步，慢指针走一步，直至快指针指向 null，此时慢节点刚好来到链表的下中节点。
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

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	count := 1
	for current != nil {
		count++

		if count == n {
			current.Next = current.Next.Next
		}

		current = current.Next
	}
	fmt.Println("判断结点:", count)
	return current
}

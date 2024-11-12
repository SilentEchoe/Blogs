package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// LeetCode.160 相交链表
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
// 官方解题思路，通过哈希表保存链表的状态，然后在某一个节点做对比
// 暴力解题思路，可以循环遍历
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	for headA != nil {
		nodeMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if nodeMap[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

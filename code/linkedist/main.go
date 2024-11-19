package main

/* 常见链表的算法题*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// leetcode.160.相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeList := make(map[*ListNode]bool)
	for headA != nil {
		nodeList[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if nodeList[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// LeetCode.206.反转链表
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// LeetCode.234. 回文链表
// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false
// 输入：head = [1,2,2,1]
// 输出：true
// 解题思路：先查到链表的中间，然后再翻转进行对比
// 简单的方式是直接用切片存储
func isPalindrome(head *ListNode) bool {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	n := len(res)
	for i, v := range res[:n/2] {
		// 最后一个元素不等于第一个元素，那么证明不是回文
		if v != res[n-i-1] {
			return false
		}
	}
	return true
}

// 实现链表中间节点有两种，一种是使用切片，一种是使用快慢指针
// 这里实际上要返回两个链表，一个是前半部分，以及后半部分
func middleNode(head *ListNode) *ListNode {
	res := []*ListNode{}
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	return res[len(res)/2]
}

// LeetCode.141. 环形链表
// 给你一个链表的头节点 head ，判断链表中是否有环。
// head = [3,2,0,-4], pos = 1
// 输出：true
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// LeetCode.142. 环形链表2
// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// head = [3,2,0,-4], pos = 1
// 返回索引为 1 的链表节点
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			// 调整到最初到样子
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// LeetCode.21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tmp := &ListNode{-1, nil}
	p := tmp
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}

		p = p.Next
	}

	// 这里是如果将没有添加到P
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return tmp.Next
}

// LeetCode.2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{Val: 0}
	cursor := root
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		l1val := 0
		l2val := 0
		if l1 != nil {
			l1val = l1.Val
		}
		if l2 != nil {
			l2val = l2.Val
		}
		sumVal := l1val + l2val + carry
		carry = sumVal / 10
		// 如果不能被膜
		sumNode := &ListNode{
			Val: sumVal % 10,
		}
		cursor.Next = sumNode
		cursor = sumNode

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
	return root.Next
}

// LeetCode.

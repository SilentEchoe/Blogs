package main

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

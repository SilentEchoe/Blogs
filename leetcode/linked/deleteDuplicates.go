package linked

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	deleteDuplicates(head)
	return head
}

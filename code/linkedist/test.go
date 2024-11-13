package main

func main() {

}

func (l *ListNode) AddNode(val int) {
	for l != nil {
		if l.Next == nil {
			l.Next = &ListNode{Val: val, Next: nil}
		}
		l = l.Next
	}
}

func (l *ListNode) DelNode(val int) *ListNode {
	if l == nil || l.Next == nil {
		return nil
	}
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Val == val {
			return p
		}
	}
	return nil
}

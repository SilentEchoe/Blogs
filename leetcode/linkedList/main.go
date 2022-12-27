/*
	链表算法题

*/

package main

func main() {

}

// Code.21 输入两个有序链表，合并成一个新的有序链表
func mergeTwoLists(l1 ListNode, l2 ListNode) {
	//声明一个虚拟头结点
	var dummy = ListNode{-1, nil}
	p := dummy
	p1 := l1
	p2 := l2

	for p1 != (ListNode{}) && p2 != (ListNode{}) {
		// 比较 p1 和 p2 两个指针
		// 将较小的节点接到 p 指针
		if p1.Val > p2.Val {
			p.Next = &p2
			p2 = *p2.Next
		} else {
			p.Next = &p1
			p1 = *p1.Next
		}
		// P 指针继续前进
		p = *p.Next
	}

	if p1 != (ListNode{}) {
		p.Next = &p1
	}

	if p2 != (ListNode{}) {
		p.Next = &p2
	}

}

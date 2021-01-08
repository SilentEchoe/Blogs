package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nowNode := l1
	for {
		if nowNode != nil {
			// 打印节点值
			println(nowNode.Val)
			// 获取下一个节点
			nowNode = nowNode.Next
		}

		// 如果下一个节点为空，表示链表结束了
		break
	}
	return nil
}

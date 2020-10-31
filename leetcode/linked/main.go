package linked

import "fmt"

// 双向链表

type node struct {
	LastNode *node
	NextNode *node
	Value    int
}

func init() {
	nodeLink := new(node)
	nodeLink.Value = 2

	node1 := new(node)
	node1.Value = 4
	node1.LastNode = nodeLink
	nodeLink.NextNode = node1

	node2 := new(node)
	node2.Value = 3
	node2.LastNode = node1
	node1.NextNode = node2

	nowNode := nodeLink
	for {
		if nowNode != nil {
			fmt.Println(nowNode.Value)
			nowNode = nowNode.NextNode
		} else {
			break
		}
	}

}

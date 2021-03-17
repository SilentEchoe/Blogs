package main

type TreeNodes struct {
	value int
	left  *TreeNodes
	right *TreeNodes
}

func main() {
	treeNodethree := &TreeNodes{
		3,
		nil,
		nil,
	}

	treeNodeTwo := &TreeNodes{
		2,
		nil,
		nil,
	}

	treeNodeOne := &TreeNodes{
		1,
		treeNodeTwo,
		treeNodethree,
	}

	treeCount := count(treeNodeOne)
	println(treeCount)
	//invertTree(treeNodeOne)
}

// 翻转二叉树
func invertTree(root *TreeNodes) {

}

// 计算树的结点
func count(root *TreeNodes) int {
	if root == nil {
		return 0
	}
	return 1 + count(root.left) + count(root.right)
}

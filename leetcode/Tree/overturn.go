package main

type TreeNodes struct {
	value int
	left  *TreeNodes
	right *TreeNodes
}

//
//               4
//             /   \
//            2     7
//           / \   / \
//          1   3 6   9

var quequ = []int{}

func main() {

	treeNodeSix := &TreeNodes{
		6,
		nil,
		nil,
	}

	treeNodeNine := &TreeNodes{
		9,
		nil,
		nil,
	}

	treeNodeOne := &TreeNodes{
		1,
		nil,
		nil,
	}

	treeNodethree := &TreeNodes{
		3,
		nil,
		nil,
	}

	treeNodeSeven := &TreeNodes{
		7,
		treeNodeSix,
		treeNodeNine,
	}

	treeNodeTwo := &TreeNodes{
		2,
		treeNodeOne,
		treeNodethree,
	}

	// 根结点
	treeNodeFrou := &TreeNodes{
		4,
		treeNodeTwo,
		treeNodeSeven,
	}

	//treeCount := count(treeNodeFrou)
	//println(treeCount)
	//newTree := invertTree(treeNodeFrou)
	//
	//newone := newTree
	//for {
	//	if newone == nil {
	//		break
	//	}
	//	println(newone.value)
	//	newone = newone.left
	//}

	binaryTreeNnfold(treeNodeFrou)
	for _, v := range quequ {
		println(v)
	}
}

// 翻转二叉树
func invertTree(root *TreeNodes) *TreeNodes {
	if root == nil {
		return nil
	}

	var tmp = root.left
	root.left = root.right
	root.right = tmp

	invertTree(root.left)
	invertTree(root.right)

	return root

}

// 计算树的结点
func count(root *TreeNodes) int {
	if root == nil {
		return 0
	}
	return 1 + count(root.left) + count(root.right)
}

// 展开二叉树成一条链表
func binaryTreeNnfold(root *TreeNodes) {
	if root == nil {
		return
	}
	quequ = append(quequ, root.value)
	binaryTreeNnfold(root.left)
	binaryTreeNnfold(root.right)
}

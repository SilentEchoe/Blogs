package main

type TreeNodes struct {
	value int
	left  *TreeNodes
	right *TreeNodes
}

//											对称二叉树
//               4								1
//             /   \						  /   \
//            2     7						2		2
//           / \   / \					  /   \   /   \
//          1   3 6   9					3	   4 3		4
//         / \
//        8   8
var quequ = []int{}

func main() {

	treeNodenight := &TreeNodes{
		8,
		nil,
		nil,
	}

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
		treeNodenight,
		treeNodenight,
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

	treeNodesThree := &TreeNodes{
		3,
		nil,
		nil,
	}

	treeNodesFrou := &TreeNodes{
		4,
		nil,
		nil,
	}

	treeNodesTwoLeft := &TreeNodes{
		2,
		treeNodesThree,
		treeNodesFrou,
	}
	treeNodesTwoRight := &TreeNodes{
		2,
		treeNodesFrou,
		treeNodesThree,
	}

	treeNodesOne := &TreeNodes{
		1,
		treeNodesTwoLeft,
		treeNodesTwoRight,
	}
	println(isSymmetric(treeNodesOne))

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

	//binaryTreeNnfold(treeNodeFrou)
	//for _, v := range quequ {
	//	println(v)
	//}

	println(TreeCount(treeNodeFrou))
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

func TreeCount(root *TreeNodes) int {
	if root == nil {
		return 0
	}

	leftCount := 1 + TreeCount(root.left)
	rightCount := 1 + TreeCount(root.right)
	if leftCount > rightCount {
		return leftCount
	}
	return rightCount
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

// 判断是否为平衡二叉树
func isSymmetric(root *TreeNodes) bool {
	if root == nil {
		return true
	}

	return NewisSymmetric(root.left, root.right)
}

// 判断是否为平衡二叉树
func NewisSymmetric(node1 *TreeNodes, node2 *TreeNodes) bool {
	// 如果都为空，代表为最底层结点
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.value == node2.value {
		return true
	}
	return NewisSymmetric(node1.left, node2.right) && NewisSymmetric(node1.right, node2.left)
}

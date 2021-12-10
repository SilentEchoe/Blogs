package main

/* 二叉搜索树 */

// BST 的每一个结点,左子树结点的值都比 node 的值要小,右子树的值都比 node 的值大
// BST 的每一个结点，它的左侧子树和右侧子树都是 BST

func main() {

}

// BST 的中序遍历结果是一个有序数组（升序）
func kthSmallest(root *TreeNode, k int) int {
	traverseBST(root, k)
	return res
}

var res, rank = -1, 0

func traverseBST(root *TreeNode, k int) {
	if root == nil || res != -1 {
		return
	}
	traverseBST(root.Left, k)
	// 中序遍历代码位置
	rank++
	if rank == k {
		res = root.Val
	}
	traverseBST(root.Right, k)
}

package main

/* 二叉搜索树 */

// BST 的每一个结点,左子树结点的值都比 node 的值要小,右子树的值都比 node 的值大
// BST 的整个左子树都要小于 root.val ,整个右子树都要大于 root.val
// BST 的每一个结点，它的左侧子树和右侧子树都是 BST
// BST 相关的问题,要么利用 BST 左小右大的特性提升算法效率,要么利用中序遍历的特性满足题目的需求
func main() {

}

// BST 的中序遍历结果是一个有序数组（升序）
// 如果想要降序打印结点的值,只需要把递归的顺序更改一下: traverse(right) -> 打印 -> traverse(left)
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

// BST 转化为累加树
func convertBST(root *TreeNode) *TreeNode {
	traverseConvertBST(root)
	return root
}

var sum = 0

func traverseConvertBST(root *TreeNode) {
	if root == nil {
		return
	}

	traverseConvertBST(root.Right)
	sum += root.Val
	root.Val = sum
	traverseConvertBST(root.Left)
}

func isValidBST(root *TreeNode) bool {
	return newIsValidBST(root, nil, nil)
}

func newIsValidBST(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return false
	}
	// 使用 max 和 min 限制 BST的结点大小
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return newIsValidBST(root.Left, min, root) && newIsValidBST(root.Right, root, max)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 获取右子树的最小结点
		var minNode = getMin(root.Right)
		root.Right = deleteNode(root.Right, minNode.Val)
		// 用右子树最小的结点替换 root 结点
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

package main

//记录最大深度
var res = 0

//记录遍历到的节点的深度
var depth = 0

func maxDepth(root *TreeNode) int {
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	depth++

	// 如果该节点的左右两个字节点为空，证明到达最深的叶子节点
	if root.left == nil && root.right == nil {
		if res < depth {
			res = depth
		}
	}
	traverse(root.left)
	traverse(root.right)
	depth--
}

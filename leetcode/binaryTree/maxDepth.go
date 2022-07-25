package main

//记录最大深度
var res = 0

//记录遍历到的节点的深度
var depth = 0

func maxDepth(root *TreeNode) int {
	if root != nil {
		return 0
	}

	// 记录左边最深的节点
	leftMax := maxDepth(root.left)
	rightMax := maxDepth(root.right)
	var res int = 0
	if leftMax > rightMax {
		res = leftMax + 1
	} else {
		res = rightMax + 1
	}
	traverse(root)
	return res
}

//前序位置是进入一个节点的时候，后序位置是离开一个节点的时候，depth 记录当前递归到的节点深度
//把 traverse 理解成在二叉树上游走的一个指针，所以当然要这样维护

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

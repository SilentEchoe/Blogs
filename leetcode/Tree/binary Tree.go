package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var tree7 = TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	var tree15 = TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}

	var tree20 = TreeNode{
		Val:   20,
		Left:  &tree15,
		Right: &tree7,
	}
	var tree9 = TreeNode{
		Val:   9,
		Left:  &tree15,
		Right: &tree7,
	}

	var tree3 = TreeNode{
		Val:   3,
		Left:  &tree9,
		Right: &tree20,
	}

	//println("参数为%d:", preorderTraversal(&tree3))
	//for _, v := range preorderTraversal(&tree3) {
	//	println(v)
	//}

	var i = maxDepth(&tree3)
	println(i)
}

func maxDepth(root *TreeNode) int {
	// 返回条件处理
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	// conquer：合并左右子树结果
	if left > right {
		return left + 1
	}
	return right + 1
}

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

	//var i = maxDepth(&tree3)
	//println(i)

	var req = divideAndConquer(&tree3)
	println(req)

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

func divideAndConquer(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// divide：分左右子树分别计算
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}

	// conquer：合并左右子树结果
	if left > right {
		return left + 1
	}
	return right + 1
}

// 先序遍历
func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}

	println(node.Val)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

// 后序遍历
func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PreOrder(node.Left)
	PreOrder(node.Right)
	println(node.Val)
}

// 中序遍历
func MidOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PreOrder(node.Left)
	println(node.Val)
	PreOrder(node.Right)
}

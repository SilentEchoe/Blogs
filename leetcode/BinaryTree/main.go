package main

// 二叉树算法题
// 	   4
//   /   \
//  2     7
// / \   / \
//1   3 6   9

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	var one = TreeNode{
		Val: 1,
		Left : nil,
		Right : nil,
	}

	var three = TreeNode{
		Val: 3,
		Left : nil,
		Right : nil,
	}

	var two = TreeNode{
		Val: 2,
		Left : &one,
		Right : &three,
	}

	var six = TreeNode{
		Val: 6,
		Left : nil,
		Right : nil,
	}

	var nine = TreeNode{
		Val: 9,
		Left : nil,
		Right : nil,
	}

	var seven = TreeNode{
		Val: 7,
		Left : &six,
		Right : &nine,
	}

	var head = TreeNode{
		Val: 4,
		Left : &two,
		Right : &seven,
	}

	mirrorTree(&head)
}

//输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
		return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// 为什么返回-1呢？（变量具有二义性）
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

// 输入一个二叉树,反转它
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Left = right
	root.Right = left
	return  root
}


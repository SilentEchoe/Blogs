package main

// 二叉树算法框架

type TreeNode struct {
	val   interface{}
	left  *TreeNode
	right *TreeNode
}

// 所谓的 traverse 函数就是一个能遍历所有节点的一个函数，与遍历数组或者链表没有本质区别
//前序位置的代码在刚刚进入一个二叉树节点的时候执行；
//后序位置的代码在将要离开一个二叉树节点的时候执行；
//中序位置的代码在一个二叉树节点左子树都遍历完，即将开始遍历右子树的时候执行。
func traverseDemo(root *TreeNode) {
	if root == nil {
		return
	}
	// 前序位置
	traverse(root.left)
	// 中序位置
	traverse(root.right)
	// 后序位置
}

func main() {

}

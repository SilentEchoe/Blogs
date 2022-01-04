package main

import (
	"fmt"
)

type TreeNodeBST struct {
	Val   int
	Left  *TreeNodeBST
	Right *TreeNodeBST
}

/* 二叉搜索树 */

// BST 的每一个结点,左子树结点的值都比 node 的值要小,右子树的值都比 node 的值大
// BST 的整个左子树都要小于 root.val ,整个右子树都要大于 root.val
// BST 的每一个结点，它的左侧子树和右侧子树都是 BST
// BST 相关的问题,要么利用 BST 左小右大的特性提升算法效率,要么利用中序遍历的特性满足题目的需求

var NodeFive = TreeNodeBST{
	Val:   5,
	Left:  &NodeThree,
	Right: &NodeSix,
}
var NodeThree = TreeNodeBST{
	Val:   3,
	Left:  &NodeTwo,
	Right: &NodeFour,
}
var NodeSix = TreeNodeBST{
	Val:   6,
	Left:  nil,
	Right: nil,
}
var NodeTwo = TreeNodeBST{
	Val:   2,
	Left:  &NodeOne,
	Right: nil,
}
var NodeFour = TreeNodeBST{
	Val:   4,
	Left:  nil,
	Right: nil,
}
var NodeOne = TreeNodeBST{
	Val:   1,
	Left:  nil,
	Right: nil,
}

func main() {

	fmt.Println(GetMinBST(&NodeFive))
	fmt.Println(GetMaxBST(&NodeFive))
}

// BST 的中序遍历结果是一个有序数组（升序）
// 如果想要降序打印结点的值,只需要把递归的顺序更改一下: traverse(right) -> 打印 -> traverse(left)
func kthSmallest(root *TreeNodeBST, k int) int {
	traverseBST(root, k)
	return res
}

var res, rank = -1, 0

func traverseBST(root *TreeNodeBST, k int) {
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
func convertBST(root *TreeNodeBST) *TreeNodeBST {
	traverseConvertBST(root)
	return root
}

var sum = 0

func traverseConvertBST(root *TreeNodeBST) {
	if root == nil {
		return
	}

	traverseConvertBST(root.Right)
	sum += root.Val
	root.Val = sum
	traverseConvertBST(root.Left)
}

func isValidBST(root *TreeNodeBST) bool {
	return newIsValidBST(root, nil, nil)
}

func newIsValidBST(root *TreeNodeBST, min *TreeNodeBST, max *TreeNodeBST) bool {
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

func insertIntoBST(root *TreeNodeBST, val int) *TreeNodeBST {
	if root == nil {
		return &TreeNodeBST{
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

func deleteNode(root *TreeNodeBST, key int) *TreeNodeBST {
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

func getMin(root *TreeNodeBST) *TreeNodeBST {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// 思路：
// 用左树 * 右树
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

// 后序遍历
// 如果当前节点要做的事情需要通过左右子树的计算结果推导出来，就要用到后序遍历

var maxSum = 0

func maxSumBST(root *TreeNodeBST) {

}

// 函数返回 int[4]{isBST,min,max,sum}
// isBST 为1则是BST,为0则不是
// min root 为根的二叉树所有节点中的最小值
// max root 为根的二叉树所有节点中的最大值
// root 为根的二叉树所有节点的和
func traverseMaxSum(root *TreeNodeBST) [4]int {
	if root == nil {
		return [4]int{1, 0, 0, 0}
	}

	var left = traverseMaxSum(root.Left)
	var right = traverseMaxSum(root.Right)

	// 后序遍历位置
	var res = [4]int{}
	// 左右节点必须是 BST,并且当前的结点值大于左树节点中最大值,小于右树的最小结点值
	if left[0] == 1 && right[0] == 1 && root.Val > left[2] && root.Val < right[1] {
		res[0] = 1
		// 计算 root 的最小值
		res[1] = GetMinBST(root)
		// 计算 root
		res[2] = GetMaxBST(root)
		// 计算总和
		res[3] = left[3] + right[3] + root.Val
	} else {
		res[0] = 0
	}
	return res
}

// 获取这棵BST树的最小值
// 使用BST 中序遍历
var minBST = []int{}

func GetMinBST(root *TreeNodeBST) int {
	if root == nil {
		return 0
	}
	GetMinBST(root.Left)
	minBST = append(minBST, root.Val)
	GetMinBST(root.Right)
	if len(minBST) > 0 {
		return minBST[0]
	}
	return 0
}

var maxBST = []int{}

func GetMaxBST(root *TreeNodeBST) int {
	if root == nil {
		return 0
	}
	GetMinBST(root.Left)
	minBST = append(minBST, root.Val)
	GetMinBST(root.Right)
	if len(minBST) > 0 {
		return minBST[len(minBST)-1]
	}
	return 0
}

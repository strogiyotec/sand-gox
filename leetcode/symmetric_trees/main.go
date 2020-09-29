package main

import "fmt"

func main() {
	node := TreeNode{Val: 1}
	node.Left = &TreeNode{Val: 2}
	node.Right = &TreeNode{Val: 2}
	node.Right.Left = &TreeNode{Val: 4}
	node.Right.Right = &TreeNode{Val: 3}
	node.Left.Left = &TreeNode{Val: 3}
	node.Left.Right = &TreeNode{Val: 4}
	fmt.Println(isSymmetric(&node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return areEqual(root.Left, root.Right)
}
func areEqual(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return areEqual(left.Left, right.Right) && areEqual(left.Right, right.Left)
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := &TreeNode{Val: 2}
	node.Left = &TreeNode{Val: 1}
	node.Right = &TreeNode{Val: 3}
	fmt.Println(isValidBST(node))
}

func isValidBST(root *TreeNode) bool {
	if root != nil {
		if root.Left != nil && root.Val < root.Left.Val {
			return false
		}

		if root.Right != nil && root.Val > root.Right.Val {
			return false
		}
		return isValidBST(root.Left) && isValidBST(root.Right)
	}
	return true
}

package main

import "fmt"

func main() {
	node := TreeNode{Val: 1}
	node.Left = &TreeNode{Val: 2}
	node.Right = &TreeNode{Val: 3}
	fmt.Println(maxPathSum(&node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := 0
	findMax(&max, root)
	return max
}

func findMax(max *int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := findMax(max, root.Left)
	left = maxOf(left, 0)
	right := findMax(max, root.Right)
	right = maxOf(right, 0)
	sum := left + right + root.Val
	*max = maxOf(*max, sum)
	return maxOf(left, right) + root.Val
}

func maxOf(first int, last int) int {
	if first > last {
		return first
	}
	return last
}

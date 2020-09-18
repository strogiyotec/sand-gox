package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := &TreeNode{Val: 3}
	left := &TreeNode{Val: 1}
	node.Left = left
	leftRight := &TreeNode{Val: 2}
	left.Right = leftRight
	right := &TreeNode{Val: 4}
	node.Right = right
	fmt.Println(kthSmallest(node, 1))
}

func kthSmallest(root *TreeNode, k int) int {
	array := make([]int, k)
	position := 0
	findSmallest(array, root, &position)
	return array[len(array)-1]
}

func findSmallest(array []int, root *TreeNode, position *int) {
	if root != nil {
		findSmallest(array, root.Left, position)
		if *position != len(array) {
			array[*position] = root.Val
			*position++
		}
		findSmallest(array, root.Right, position)
	}
}

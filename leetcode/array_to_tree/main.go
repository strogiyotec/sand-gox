package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return binarySearch(nums, 0, len(nums)-1)
}

func binarySearch(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	middle := left + (right-left)/2
	node := TreeNode{Val: nums[middle]}
	node.Left = binarySearch(nums, left, middle-1)
	node.Right = binarySearch(nums, middle+1, right)
	return &node

}

func main() {
	fmt.Println(sortedArrayToBST([]int{0, 1, 2, 3, 4, 5}).Val)
}

package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{2, 2}, 2))
	fmt.Println(searchRange([]int{1, 2, 3}, 1))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}
	first := first(nums, target)
	if first == -1 {
		return []int{-1, -1}
	} else {
		last := last(nums, target, first)
		return []int{first, last}
	}
}

func last(nums []int, target int, left int) int {
	right := len(nums) - 1
	index := -1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] <= target {
			if nums[middle] == target {
				index = middle
			}
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return index
}
func first(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	index := -1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] >= target {
			if nums[middle] == target {
				index = middle
			}
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return index
}

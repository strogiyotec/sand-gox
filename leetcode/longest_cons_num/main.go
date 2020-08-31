package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

// https://leetcode.com/problems/longest-consecutive-sequence/submissions/
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	max := 0
	for _, num := range nums {
		left := num - 1
		right := num + 1
		for m[left] {
			delete(m, left)
			left--
		}
		for m[right] {
			delete(m, right)
			right++
		}
		diff := right - left - 1
		if max < diff {
			max = diff
		}
		if len(m) == 0 {
			break
		}
	}
	return max
}

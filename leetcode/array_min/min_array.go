package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(s int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	max := math.MaxInt32
	for true {
		if sum >= s {
			for sum >= s {
				if max > right-left {
					max = right - left
				}
				sum -= nums[left]
				left++
			}
		} else {
			if right < len(nums) {
				sum += nums[right]
				right++
			} else {
				break
			}
		}
	}
	if max == math.MaxInt32 {
		return 0
	} else {
		return max
	}
}

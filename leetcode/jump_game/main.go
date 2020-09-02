package main

import (
	"fmt"
)

type pair struct {
	first  int
	second int
}

func main() {
	fmt.Println(canJump([]int{0}))
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	tuple := pair{
		first:  0,
		second: 0,
	}
	for true {
		canJump := -1
		for i := tuple.first; i <= tuple.second; i++ {
			if canJump < nums[i]+i {
				canJump = nums[i] + i
			}
		}
		if canJump >= len(nums)-1 {
			return true
		}
		tuple = pair{
			first:  tuple.second + 1,
			second: canJump,
		}
		if tuple.first > tuple.second {
			return false
		}
	}
	return false
}

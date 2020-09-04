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
	if len(nums) == 1 {
		return true
	}
	index := -1
	tuple := pair{
		first:  0,
		second: 0,
	}
	for true {
		for i := tuple.first; i <= tuple.second; i++ {
			if i+nums[i] > index {
				index = i + nums[i]
			}
		}
		if index >= len(nums)-1 {
			return true
		}
		tuple = pair{
			first:  tuple.second,
			second: index,
		}
		if tuple.first > tuple.second {
			return false
		}
	}
	return false
}

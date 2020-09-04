package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	sum := 0
	max := 0
	for index, value := range nums {
		if index == 0 {
			sum = value
			max = sum
		} else {
			diff := value + sum
			if value > diff {
				sum = value
			} else {
				sum = diff
			}
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

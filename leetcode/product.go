package main

import "fmt"

func main() {
	nums := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(nums)
	nums = productExceptSelf([]int{1, 0})
	fmt.Println(nums)
	nums = productExceptSelf([]int{0, 0})
	fmt.Println(nums)
	nums = productExceptSelf([]int{0, 4, 0})
	fmt.Println(nums)
}

func productExceptSelf(nums []int) []int {
	first := make([]int, len(nums))
	second := make([]int, len(nums))
	first[0] = 1
	second[len(second)-1] = 1
	for i := 0; i < len(nums)-1; i++ {
		first[i+1] = first[i] * nums[i]
	}
	// fmt.Println(first)
	for i := len(nums) - 1; i > 0; i-- {
		second[i-1] = second[i] * nums[i]
	}

	// fmt.Println(second)
	for i := 0; i < len(nums); i++ {
		nums[i] = first[i] * second[i]
	}
	return nums
}

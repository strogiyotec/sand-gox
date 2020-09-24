package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	solution := [][]int{}
	comb := []int{}
	return permutations(&nums, 0, len(nums)-1, &solution, &comb)
}

func permutations(nums *[]int, from int, end int, solution *[][]int, permutation *[]int) {
	if from == end {
		append(solution, permutation)
	}
	for i := from; i <= end; i++ {
		swap(nums, from, i)
		permutations(nums, from+1, end, solution, permutation)
		swap(nums, from, i)
	}
}

func swap(nums *[]int, from int, to int) {
	temp := from
	nums[from] = to
	nums[to] = temp
}

package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	solution := [][]int{}
	comb := []int{}
	index := 0
	permutations(&nums, 0, len(nums)-1, &solution, &comb, &index)
	return solution
}

func permutations(nums *[]int, from int, end int, solution *[][]int, permutation *[]int, index *int) {
	if from == end {
		tmp := make([]int, len(*nums))
		copy(tmp, *nums)
		*solution = append(*solution, tmp)
	}
	for i := from; i <= end; i++ {
		swap(nums, from, i)
		permutations(nums, from+1, end, solution, permutation, index)
		swap(nums, from, i)
	}
}

func swap(nums *[]int, from int, to int) {
	temp := (*nums)[from]
	(*nums)[from] = (*nums)[to]
	(*nums)[to] = temp
}

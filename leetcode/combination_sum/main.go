package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	slice := [][]int{}
	currentCandidates := []int{}
	calculate(candidates, target, &slice, currentCandidates, 0)
	return slice
}

func calculate(candidates []int, target int, slice *[][]int, currentCandidates []int, position int) bool {
	if target < 0 {
		return false
	}
	if target == 0 {
		temp := make([]int, len(currentCandidates))
		copy(temp, currentCandidates)
		*slice = append(*slice, temp)
	}
	for i := position; i < len(candidates); i++ {
		currentCandidates = append(currentCandidates, candidates[i])
		calculate(candidates, target-candidates[i], slice, currentCandidates, i)
		currentCandidates = currentCandidates[:len(currentCandidates)-1]
	}
	return true
}

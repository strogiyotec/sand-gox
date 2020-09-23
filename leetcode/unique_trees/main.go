package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

type intPair struct {
	first  int
	second int
}

func numTrees(n int) int {
	cache := make(map[intPair]int)
	return find(1, n, n, &cache)
}
func find(start int, end int, limit int, storage *map[intPair]int) int {
	if start < 0 || start >= end || end > limit {
		return 1
	}
	pair := intPair{first: start, second: end}
	if val, ok := (*storage)[pair]; ok {
		return val
	}
	cnt := 0
	for i := start; i <= end; i++ {
		left := find(start, i-1, limit, storage)
		right := find(i+1, end, limit, storage)
		cnt += left * right
	}
	if start == 1 && end == 2 {
		fmt.Printf("Salam %d \n", cnt)
	}
	(*storage)[pair] = cnt
	return cnt
}

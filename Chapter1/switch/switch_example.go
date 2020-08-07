package main

import "fmt"

func main() {
	fmt.Println(Sign(10))
}

func Sign(x int) int {
	switch {
	case x > 0:
		return 1
	default:
		return 0
	case x < 0:
		return -1
	}
}

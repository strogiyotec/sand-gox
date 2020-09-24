package main

import "fmt"

func main() {
	var first int
	first = 1
	var second int
	second = 2
	fmt.Printf("%d %d \n", first, second)
	swap(&first, &second)
	fmt.Printf("%d %d \n", first, second)
}

func swap(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp

}

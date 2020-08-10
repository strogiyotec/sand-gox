package main

import "fmt"

func main() {
	x, y := 2, 3
	fmt.Printf("%d %d \n", x, y)
	x, y = y, x
	fmt.Printf("%d %d ", x, y)
}

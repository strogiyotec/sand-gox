package main

import (
	"fmt"
)

func main() {
	dynamicSize()
	dynamicSize2()
	array := []int{1, 2, 3, 4, 5, 6}
	expandSlice(array[0:4])
}

func dynamicSize() {
	names := [...]string{"Almas", "Abdrazak"}
	fmt.Println(names)
}

func dynamicSize2() {
	names := []string{"Almas", "Abdrazak"}
	fmt.Println(names)
}

func expandSlice(slice []int) {
	newOne := append(slice, 5)
	fmt.Println(slice)
	fmt.Println(newOne)
}

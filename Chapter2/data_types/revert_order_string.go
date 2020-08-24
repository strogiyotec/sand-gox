package main

import "fmt"

func main() {
	words := []string{"orange", "apple", "banana"}
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Println(words[i])
	}
}

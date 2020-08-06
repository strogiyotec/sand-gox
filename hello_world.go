package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("Index %d and word is %s\n", i, arg)
	}
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	// read input
	input := bufio.NewScanner(os.Stdin)
	wordCount := make(map[string]int)
	for input.Scan() {
		wordCount[input.Text()]++
	}
	for word, count := range wordCount {
		fmt.Printf("Word %s , appearances %d\n", word, count)
	}
}

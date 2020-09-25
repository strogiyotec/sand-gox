package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}

func uncommonFromSentences(A string, B string) []string {
	counter := make(map[string]int)
	for _, word := range strings.Split(A, " ") {
		if val, ok := counter[word]; ok {
			val++
			counter[word] = val
		} else {
			counter[word] = 1
		}
	}
	for _, word := range strings.Split(B, " ") {
		if val, ok := counter[word]; ok {
			val++
			counter[word] = val
		} else {
			counter[word] = 1
		}
	}
	keys := make([]string, 0, len(counter))
	for k, v := range counter {
		if v == 1 {
			keys = append(keys, k)
		}
	}
	return keys

}

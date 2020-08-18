package main

import (
	"fmt"
	"test_package/simpleinterest"
)

// init from simpleinterest will be called first
func init() {
	fmt.Println("I am main")
}

func main() {
	fmt.Println(simpleinterest.Calculate(2, 5))
}

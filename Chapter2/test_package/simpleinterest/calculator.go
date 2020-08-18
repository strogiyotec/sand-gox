package simpleinterest

import "fmt"

// called only once when package is imported
func init() {
	fmt.Println("Init was called")
}

// capital leter means that this method is public
func Calculate(first int, second int) int {
	return first + second
}

// this method is private
func calculate(first int, second int) int {
	return first + second
}

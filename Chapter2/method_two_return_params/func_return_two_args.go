package main

import (
	"fmt"
	"strings"
)

func main() {
	toUpper, failed := upperCase("orange")
	if !failed {
		fmt.Println("Nope")
	} else {
		fmt.Println(toUpper)
	}
}

func upperCase(name string) (string, bool) {
	if strings.Compare(name, "orange") == 0 {
		return strings.ToUpper(name), true
	} else {
		return name, false
	}
}

package main

import (
	"flag"
	"fmt"
	"strings"
)

var test = flag.Bool("test", false, "Provide a valid test")

func main() {
	flag.Parse()
	if !*test {
		fmt.Println("Wrong")
	} else {
		fmt.Println(strings.Join(flag.Args(), ", "))
	}
}

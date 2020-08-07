package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[2:] {
		fmt.Println(arg)
	}
}

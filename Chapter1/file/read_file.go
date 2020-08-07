package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, e := ioutil.ReadFile("names.txt")
	if e != nil {
		fmt.Errorf("Error is %s", e)
	} else {
		for _, word := range strings.Split(string(data), "\n") {
			fmt.Println(word)
		}
	}
}

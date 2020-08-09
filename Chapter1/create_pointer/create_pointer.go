package main

import "fmt"

func main() {
	//create pointer to int without name(p is a name of pointer not a variable)
	p := new(int)
	*p = 5
	fmt.Println(*p)
}

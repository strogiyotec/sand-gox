package main

import "fmt"

func main() {
	x := 2
	p := &x
	fmt.Println(*p)
	*p = 4
	fmt.Println(x)
	fmt.Println("Increment example")
	increment(&x)
	//was 4 became 5
	fmt.Println(x)

}

func increment(p *int) {
	*p++
}

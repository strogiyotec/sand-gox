package main

import "fmt"

type Celsius float64

func main() {
	printIt()
}

func printIt() {
	temperature := Celsius(21.2)
	fmt.Println(temperature)
}

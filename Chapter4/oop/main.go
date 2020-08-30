package main

import (
	"fmt"
	"oop/emp"
)

func main() {
	old := emp.OldEmp()
	fmt.Println(old.Age)
	e := emp.New("Almas", "Abdrazak", 21)
	fmt.Println(e)
}

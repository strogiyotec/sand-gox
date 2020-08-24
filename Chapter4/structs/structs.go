package main

import "fmt"

type programmer struct {
	lazy   bool
	skills []string
}

func main() {
	senior := lazySenior()
	modify(senior)
	//still true
	fmt.Println(senior)
	//now it's false
	modifyPointer(&senior)
	fmt.Println(senior)
}

func lazySenior() programmer {
	return programmer{
		lazy:   true,
		skills: []string{"CRUD"},
	}
}

func junior() programmer {
	return programmer{
		lazy: false,
	}
}

func modify(emp programmer) {
	emp.lazy = !emp.lazy
}
func modifyPointer(emp *programmer) {
	emp.lazy = !emp.lazy
}

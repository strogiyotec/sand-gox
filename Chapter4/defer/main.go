package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	array := readAll("main.go")
	for _, element := range array {
		fmt.Println(element)
	}
}

func readAll(path string) []string {
	file, e := os.Open(path)
	if e != nil {
		return []string{}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	array := make([]string, 10)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	return array

}

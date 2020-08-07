package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, er := http.Get("https://example.com")
	if er != nil {
		fmt.Println(er)
		return
	} else {
		if er != nil {
			fmt.Println(er)
			return
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Header)
			response.Body.Close()
		}
	}
}

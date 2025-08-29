package main

import (
	"fmt"
	"strings"
)

func main() {

	var partial string

	fmt.Print("Enter a search string: ")
	fmt.Scanln(&partial)

	fmt.Println("Results:")

	for i := range books {
		if strings.Contains(strings.ToLower(books[i]), strings.ToLower(partial)) {
			fmt.Println(books[i])
		}
	}
}

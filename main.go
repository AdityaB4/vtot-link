package main

import "fmt"

func main() {
	n, err := fmt.Printf("hello")
	if err != nil {
		return
	}
	fmt.Printf("number of bytes: %v\n", n)
}

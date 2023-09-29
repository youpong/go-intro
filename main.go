package main

import "fmt"

func inValid() bool {
	return true
}

func main() {
	if inValid() {
		fmt.Println("!")
	}
	fmt.Println("Hello")
}

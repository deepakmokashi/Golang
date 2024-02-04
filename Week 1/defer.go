package main

import (
	"fmt"
)

func main() {

	//normal case
	fmt.Println("Without defer keyword")
	fmt.Println("Hello")
	fmt.Println("World")

	//using defer
	fmt.Println("With defer keyword")
	defer fmt.Println("Hello")
	fmt.Println("World")
}
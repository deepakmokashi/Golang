package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Default value of ptr: ", ptr)

    i:=10
	ptr = &i

	fmt.Println("Address of pointer", ptr)
	fmt.Println("Value of pointer", *ptr)

	*ptr = i + 2

	fmt.Println("After addition", i)
}
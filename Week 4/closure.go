package main

import "fmt"

func main() {
	i:= 0

	count := func () int {
		i++;
		return i;
	}

	fmt.Println("Value of i",count())
	fmt.Println("Value of i",count())
}
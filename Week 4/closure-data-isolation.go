package main

import "fmt"

//new count function
func newCount() func() int {
	i:=10	
	return func() int {
		i++;
		return i
	}
}

func main() {
	count := newCount()

	fmt.Println(count())
	fmt.Println(count())
}
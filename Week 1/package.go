package main

import (
	"fmt"
	"math"
)

func add( x int, y int) int {
	return x + y;
}
func sub( x , y int) int {
	return x - y;
}
func main() {
	fmt.Println(math.Pi)
	fmt.Println(add(5,6))
	fmt.Println(sub(10,2))
}
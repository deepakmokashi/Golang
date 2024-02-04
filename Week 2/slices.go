package main

import "fmt"

func main() {
	prime:=[]int {2,3,5,7,11}

	fmt.Println("Printing slice", prime)

	a:= prime[0:3]
	b:= prime[2:5]

	fmt.Println("Printing new values:",a,b)
}
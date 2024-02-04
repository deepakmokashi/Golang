package main

import "fmt"

func main() {
	a:= make([]int,5)
	fmt.Println(a," Length",len(a)," Capacity ",cap(a))

	b:= make([]int,0,5)
	fmt.Println(b," Length ",len(b)," Capacity ",cap(b))

	c:= b[:2]
	fmt.Println(c," Length ",len(c)," Capacity ",cap(c))

	d:= c[2:3]
	fmt.Println(d," Length ",len(d)," Capacity ",cap(d))
}
package main

import "fmt"

func main () {
	var s [] int
	fmt.Println(s," Length",len(s)," Capacity ",cap(s))

	s = append(s,1)
	fmt.Println(s," Length",len(s)," Capacity ",cap(s))

	s = append(s, 2, 3, 4)
	fmt.Println(s," Length",len(s)," Capacity ",cap(s))
	

}
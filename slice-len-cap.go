package main

import "fmt"

func main() {
	even:=[] int {0,2,4,5,8}
	fmt.Println(even)
	fmt.Println("length",len(even), "capacity",cap(even))

	even = even[:2]
	fmt.Println("length",len(even), "capacity",cap(even))

	even = even[2:]
	fmt.Println("length",len(even), "capacity",cap(even))
}
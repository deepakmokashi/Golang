package main

import "fmt"

func main() {
	ElctronicDevics:= make(map[int]string)

	ElctronicDevics[1] = "Mobile"

	fmt.Println("Value:",ElctronicDevics[1])

	ElctronicDevics[1] = "Laptop"
	fmt.Println("Value:",ElctronicDevics[1])

	delete(ElctronicDevics,1)
	fmt.Println("Value:",ElctronicDevics[1])

	v,ok:= ElctronicDevics[1]
	fmt.Println("The value",v,"Presnet?",ok)

}
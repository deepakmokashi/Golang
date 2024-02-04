package main

import "fmt"

type Address struct {
	City string
	Pincode int
}

func main() {
	a1 := Address{"Pune",411001}
	fmt.Println("Address1:",a1)

	fmt.Println("City:",a1.City)
	fmt.Println("Pincode:",a1.Pincode)

	a1.City = "Mumbai"

	fmt.Println("Address :",a1)

	
}
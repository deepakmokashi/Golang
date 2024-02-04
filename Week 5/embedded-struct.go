package main

import "fmt"

type Address struct {
	City string
	Pin int
}

type Person struct {
	Name string
	Age int
	Address //embedded struct
}

func main() {
     p:= Person{
		Name:"Jon",
		Age: 20,
		Address : Address{
			City:"Pune",
			Pin: 411111,
		},
	 }

	 fmt.Println("Details: ",p)
	 fmt.Println("Name: ",p.Name)
	 fmt.Println("City: ",p.City)
}
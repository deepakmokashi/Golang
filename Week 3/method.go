package main

import "fmt"

type Person struct {
	Name string
	Age int
	Email string
}

func main() {
	person1 := Person{"Rob",20,"rob@gmail.com"}
	fmt.Printf("Person1: %+v\n",person1)

	// person2 := Person{Name: "Mike", Age: 30, Email: "mike22@gmail.com"}
	// fmt.Printf("Person2: %+v\n",person2)

	// person3 := Person{Name: "Jack"}
	// fmt.Printf("Person3: %+v\n",person3)

	person1.getAge()
	person1.setAge(30)
}

func (p Person) getAge() {
     fmt.Println("Age of person:", p.Age)
}

func (p Person) setAge(age int) {
	p.Age = age
	fmt.Println("New Age of person:", p.Age)
}
package main

import (
	"fmt"
)

type geometry interface {
    area() float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println("Area of rectangle",g.area())
}


func main() {
    r := rect{width: 3, height: 4}

    measure(r)
}
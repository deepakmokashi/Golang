package main

import "fmt"

func main() {
	var a[5] int

	fmt.Println("Array ", a)

	a[0] = 10
	a[1] = 20
    
	fmt.Println("Length of array: ",len(a))
	fmt.Println("Array after value added", a)

	b:=[6] int {1,2,3,4,5,6}
	fmt.Println("Array b",b)

	var c[2][2] int
	for i:=0;i<2;i++ {
		for j:=0;j<2;j++ {
			c[i][j] = i * j;
		}
	}

	fmt.Println("2D array:", c)
}
package main

import "fmt"

func sums(nums ...int) {
	fmt.Print(nums," ")
	total := 0
	for _,num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sums(1,2,3)
	arr:= []int{12,13,14}
	sums(arr...)
}
package main

import "fmt"

func main() {
	nums:=[]int {2,3,5,7,11}
	sum := 0
	for _,num :=range nums{
		sum += num
	}
	fmt.Println("Sum",sum)

	for i,num := range nums{
		if num==3 {
			fmt.Println("Index",i)
		}
	}
}
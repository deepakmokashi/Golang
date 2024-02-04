package main

import "fmt"

func main() {
	var fruitlist = []string{"Mango","Banana","Apple"}

	fmt.Printf("Type of fruitlist is %T\n",fruitlist)
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist,"Peach")
	fmt.Println(fruitlist)

	fruitSeasonal := fruitlist[1:2]
	fmt.Println(fruitSeasonal)
}
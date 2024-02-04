package main

import (
	"fmt"
)

func main() {
	i:= "weekend"

	switch i {
	case "weekend":
		fmt.Println("Yeah its weekend!!!! :)")
	case "weekday":
		fmt.Println("Phir se monday :(")
	}
}
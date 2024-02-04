package main

import "fmt"

func main() {
	Mobiles:= make(map[int]string)

	Mobiles[1] = "RealMe"
	Mobiles[2] = "Oppo"
	Mobiles[3] = "Apple"
	fmt.Println(Mobiles)

	for key,value := range Mobiles{
		fmt.Printf("Key is %v, value is %v\n",key,value)
	}
}
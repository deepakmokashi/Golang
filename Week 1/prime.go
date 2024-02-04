package main

import (
	"fmt"
)

func start(num int) {
	fmt.Println("********** Mathematical Operation ***********")
	fmt.Print("Number: ")
	fmt.Println(num)
}

func checkPrime(num int) bool {
	var tmp bool = true
	for i := 2; i < num ; i++ {
		if num % i == 0 {
           tmp = false
		   return tmp;
		}	
	}
	return tmp;
}

func checkEvenOdd (num int) bool {
	if num % 2 == 0 {
		return true
	} else {
		return false
	}
}

func checkSquare (num int) int {
	return num*num;
}

func main() {
	var num int = 6
	start(num)

	//check prime number
	var prime bool = checkPrime(num)
	if prime {
		fmt.Println("The given no is Prime")
	} else {
		fmt.Println("The Given no is not prime")
	}

	//check even-odd number
	var even bool = checkEvenOdd(num)
    if even {
		fmt.Println("The given no is Even")
	} else {
		fmt.Println("The given no is Odd")
	}

	//calculate square of number
	sqr:= checkSquare(num)
	fmt.Print("Square of given no is ")
	fmt.Println(sqr)
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("IfElse in golang")

	loginCount := 9
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10{
		result = "Watch out"
	} else {
		result = "exactly 10 logins"
	}

	fmt.Println(result)

	if(9 % 2 == 0) {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

}
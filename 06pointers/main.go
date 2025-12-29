package main

import "fmt"

func main() {
	fmt.Println("welcome to class on pointers")

	// var ptr *int

	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of pointer is", ptr)
	fmt.Println("value of pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new value is:", myNumber)
}
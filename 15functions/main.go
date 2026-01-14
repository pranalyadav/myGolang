package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes := proAdder(2,4,3,5,7,8)
	fmt.Println("Pro result is: ", proRes)
}

func adder(num1 int, num2 int) int{
	return num1 + num2
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Namstey golang")
}
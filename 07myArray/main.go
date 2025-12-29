package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "grape"

	fmt.Println("the list of fruits are:", fruitList)
	fmt.Println("the length of array is:", len(fruitList))

	var vegList = [3]string{"potato", "tomato","beans"}
	fmt.Println("the veg list is:", vegList)
	fmt.Println("the length of veg array is:", len(vegList))
}
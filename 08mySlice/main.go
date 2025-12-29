package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slice in golang")

	var fruitList = []string{"apple", "banana", "grape"}
	fmt.Printf("type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "orange")
	fmt.Println("the list of fruit", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 264
	highScores[1] = 500
	highScores[2] = 300
	highScores[3] = 440
	//highScores[4] = 600 // this will throw an error as the size is only 4

	highScores = append(highScores, 550, 600, 700) // we can use append function to add more elements
	fmt.Println(highScores)

	sort.Ints(highScores) // this will sort the slice in increasing order
	fmt.Println("sorted scores are:", highScores)
}
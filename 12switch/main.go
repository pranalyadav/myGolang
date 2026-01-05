package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Switch in Golang")

	// rand.New(time.Now().UnixNano())
	diceNumber := rand.IntN(6) + 1
	fmt.Println("Value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1, you can go out")
	case 2:
		fmt.Println("You got 2, you can go out")
	case 3:
		fmt.Println("You got 3, you can go out")
		fallthrough
	case 4:
		fmt.Println("You got 4, you can go out")
	case 5:
		fmt.Println("You got 5, you can go out")
	case 6:
		fmt.Println("You got 6, and roll the dice again")
	default:
		fmt.Println("what was that?")
	}
	
}
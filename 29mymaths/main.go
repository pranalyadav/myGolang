package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("welcome to maths in golang")
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5)+1)

	//random number from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
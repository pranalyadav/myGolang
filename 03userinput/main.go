package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	//comma ok || err err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating", input)
}
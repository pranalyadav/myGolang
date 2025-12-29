package main

import "fmt"

//constant
//public variable
const LoginToken string = "gkjfk"

func main() {
	var username string = "pranal"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("variable is of type: %T \n", isloggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	// default values and aliases
	var anotherVaribale int
	fmt.Println(anotherVaribale)
	fmt.Printf("variable is of type: %T \n", anotherVaribale)
	//default value for int is 0

	var anotherVar string
	fmt.Println(anotherVar)
	fmt.Printf("variable is of type: %T \n", anotherVar)
	//default value for string is ""

	//implicit style
	var website = "learncode.in"
	fmt.Println(website)

	// no var style
	numberOfUsers := 30000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
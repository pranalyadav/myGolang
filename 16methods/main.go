package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	akshay := User{"Akshay", "akshay@example.com", true, 27}
	fmt.Println(akshay)
	fmt.Printf("Akshay details are: %+v\n", akshay)
	fmt.Printf("Name is %v and email is %v\n", akshay.Name, akshay.Email)
	akshay.GetStatus()
	akshay.NewMail()
	fmt.Printf("Name is %v and email is %v\n", akshay.Name, akshay.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of the user is:", u.Email)
}
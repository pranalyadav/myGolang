package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "python"

	fmt.Println("List of languages are", languages)
	fmt.Println("JS is short for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of languages are", languages)

	//loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files in golang")
	content := "This is sample content for the file."

	file, err := os.Create("./mycontent.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./mycontent.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside file is:", string(dataByte))
}

func checkNilError(err error) {
if err != nil {
		panic(err)
	}
}
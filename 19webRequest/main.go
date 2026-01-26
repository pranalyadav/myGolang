package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://takeuforward.org/"

func main() {
	fmt.Println("working with web request in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response of type: %T\n",response)

	defer response.Body.Close() //callers responsibility to close the body

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
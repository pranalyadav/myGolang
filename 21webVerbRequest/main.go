package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb video")
	PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("content length is:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(string(content))
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename" : "Golang",
			"price" : 0,
			"platform" : "youtube"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
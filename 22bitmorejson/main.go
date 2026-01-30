package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcomt to JSON video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course {
		{"ReactJS", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN bootcamp", 199, "learncodeonline.in", "xyz123", []string{"full-stack", "js"}},
		{"Angular bootcamp", 399, "learncodeonline.in", "pqr123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses,"", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)
}
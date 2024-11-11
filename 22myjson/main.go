package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"ReactJS Bootcamp", 100, "mylearning.in", "abc123", []string{"frontend-dev", "js"}},
		{"MERN Bootcamp", 100, "mylearning.in", "hri123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 100, "mylearning.in", "bcd123", nil},
	}

	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

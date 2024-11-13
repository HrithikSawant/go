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
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
        "coursename": "ReactJS Bootcamp",
        "Price": 100,
        "website": "mylearning.in",
        "tags": ["frontend-dev","js"]
    }
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}

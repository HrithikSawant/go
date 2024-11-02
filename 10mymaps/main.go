package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	// syntax: make(map[key]value)
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["Go"] = "Golang"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "PY")
	fmt.Println("List of all languages: ", languages)

	// loops
	for key, value := range languages {
		fmt.Printf(" For key %v, value %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf(" value is %v\n", value)
	}
}

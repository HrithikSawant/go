package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// Initializer, condition, incremental
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("Another Approach with range")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("For each")

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougueValue := 1

	for rougueValue < 20 {

		// when goto hit it will exit loop and jump to goto label
		if rougueValue == 21 {
			goto mylabel
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		if rougueValue == 16 {
			break
		}

		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

	// Even loop is exit on 16 value because of break this below label is printed
	// as it is considered to another block of code
mylabel:
	fmt.Println("Welcome to concept of goto")

}

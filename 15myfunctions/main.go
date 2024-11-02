package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proRes, proMsg := proAdder(2, 5, 6, 7, 1, 9)
	fmt.Println("ProResult is: ", proRes)
	fmt.Println("Message is: ", proMsg)
}

// function signature's
// syntax:  func name(parameter type, parameter2 type) returntype {}
func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

// when expected N number of values
// variadic functions with two returns
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}

	return total, "Hello Pro result function"
}

func greeterTwo() {
	fmt.Println("Another Function")
}

func greeter() {
	fmt.Println("Namaskaar to golang")
}

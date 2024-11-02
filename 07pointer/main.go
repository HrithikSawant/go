package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	// reference &
	myNumber := 23
	var ptr1 = &myNumber

	fmt.Println("Value of actual pointer is ", ptr1)
	fmt.Println("Value of actual pointer is ", *ptr1)

	// Performed on actual memory value not on copy of it
	*ptr1 = *ptr1 + 2
	fmt.Println("Value  is ", myNumber)
}

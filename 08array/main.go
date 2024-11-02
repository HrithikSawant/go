package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [6]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Mango"
	fruitList[3] = "Grape"
	fruitList[4] = "Banana"

	fmt.Println("Array list: ", fruitList)

	// Note len is depend on size of array not elements assign to it
	fmt.Println("length is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList)
}

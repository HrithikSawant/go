package main

import "fmt"

func linearSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	for index := 0; index < len(arr); index++ {
		//check for element if matches target
		element := arr[index]
		if element == target {
			return index
		}
	}
	return -1
}

func main() {

	arr := []int{1, -19, 20, 12, 14, 17, 3, 6, 20, -20}
	result := linearSearch(arr, 17)

	fmt.Println(result)

	result1 := linearSearch(arr, 37)
	fmt.Println(result1)
}

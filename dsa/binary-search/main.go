package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{1, -19, 20, 12, 14, 17, 3, 6, 20, -20, 19, 23}
	sort.Ints(arr) // as binary search require element to be sorted

	result := binarySearch(arr, 17)
	fmt.Println(result)

	result1 := binarySearch(arr, 37)
	fmt.Println(result1)
}

func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1 // Empty
	}
	start := 0
	end := (len(arr))

	for start < end {
		mid := (start + (end-start)/2)

		if arr[mid] == target {
			return mid
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1 // Element not found
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of GoLang")

	// present time
	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.August, 13, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
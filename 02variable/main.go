package main

import "fmt"

// Public Vaziable
const LoginToken string = "glasdglasjdfg"

func main() {
	var username string = "hrithik"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.5243572359824359872345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "hrithik.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 12341324234
	fmt.Println(numberOfUser)

	// Printing variable
	fmt.Println(LoginToken)
}

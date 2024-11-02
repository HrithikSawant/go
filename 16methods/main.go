package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	hrithik := User{"Hrithik", "hrithik@go.dev", true, 18}
	fmt.Println(hrithik)
	fmt.Printf("hrithik details are: %+v\n", hrithik)
	fmt.Printf("Name is %v and email is %v\n", hrithik.Name, hrithik.Email)

	hrithik.GetStatus()
	hrithik.NewMail()
	fmt.Printf("Name is %v and email is %v\n", hrithik.Name, hrithik.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "testing@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}

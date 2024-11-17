package main

import (
	"fmt"
	"math"
)

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

// Example: 1 Method

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods Not Struct Type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Regular Abs() with no change in functionality
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

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

	// Example1: Methods
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	// Method Not Struct Type
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

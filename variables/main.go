package main

import "fmt"

func main() {
	var name string = "John"
	fmt.Println(name)

	// type inference
	var userName = "Rahul"
	fmt.Println(userName)

	var age int = 20
	fmt.Println(age)

	// multiple variables
	var (
		firstName = "John"
		lastName  = "Doe"
		userAge   = 20
	)
	fmt.Println(firstName, lastName, userAge)

	// short declaration
	user := "Rahul"
	fmt.Println(user)

	// boolean
	var isActive bool
	isActive = true
	fmt.Println(isActive)
}

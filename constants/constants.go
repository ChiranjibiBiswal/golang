package main

import "fmt"

const age = 20

func main() {
	// constants are declared like variables, but with the const keyword
	const userName string = "John Doe"
	const userAge int = 30

	fmt.Println(userName, userAge)
	fmt.Println(age)

	// userName = "Jane Doe" // error: cannot assign to userName (constant of type string)

	// constants can be declared in a block
	const (
		host = "localhost"
		port = 5000
	)

	fmt.Println(host, port)
}

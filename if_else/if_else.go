package main

import "fmt"

func main() {
	const age int = 20

	// basic condition
	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	// logical condition check
	const userRole = "Admin"
	const hasAccess = false

	if userRole == "Admin" && hasAccess {
		fmt.Println("Authorized User.")
	} else {
		fmt.Println("Unauthorized!")
	}

	// short variable declaration with condition
	if userRole := "Retailer"; userRole == "Admin" {
		fmt.Println("User is an Admin")
	} else if userRole == "Retailer" {
		fmt.Println("User is a Retailer")
	} else {
		fmt.Println("User role is unknown")
	}

	// go does not support ternary operator, but we can use if-else to achieve similar functionality
	const isMember = true
	var discount int
	if isMember {
		discount = 10
	} else {
		discount = 0
	}
	fmt.Println("Discount:", discount)
}

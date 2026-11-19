package main

import "fmt"

// printData prints a value of any type. The type parameter T is inferred from data.
func printData[T any](data T) {
	fmt.Println("Printing data:", data)
}

// printDetails prints a value whose type can be compared.
// It demonstrates the comparable constraint, even though fmt.Println does not require comparison.
func printDetails[T comparable](data T) {
	fmt.Println("Printing data:", data)
}

func main() {
	printData("Rakesh")
	printData(24)

	printDetails("Rajesh")
	printDetails(25)
}

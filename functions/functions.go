package main

import "fmt"

// function returning single value
func addNumbers(a, b int) int {
	return a + b
}

// function returning multiple values
func multiValueReturn() (string, string, int) {
	return "John", "Doe", 25
}

// function returning another function
func printData() func(int) int {
	return func(num int) int {
		return num * num
	}
}

func main() {
	result := addNumbers(5, 10)
	fmt.Println(result)

	fmt.Println(multiValueReturn())
	firstName, _, age := multiValueReturn()
	fmt.Println(firstName, age)

	// function returning another function
	printFunc := printData()
	fmt.Println(printFunc(5))
}

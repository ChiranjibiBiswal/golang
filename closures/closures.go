package main

import "fmt"

// A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense, the function is "bound" to the variables.
func closures() func () int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}

func main() {
	increment := closures()
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
}
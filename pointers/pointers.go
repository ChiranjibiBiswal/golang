package main

import "fmt"

func passByValue(num int) {
	num = 20

	fmt.Println("Inside passByValue:", num)
}

func passByReference(num *int) {
	*num = 30

	fmt.Println("Inside passByReference: ", *num)
}

func main() {
	num := 10

	passByValue(num)
	fmt.Println("After passByValue:", num)

	passByReference(&num)
	fmt.Println("After passByReference:", num)
}
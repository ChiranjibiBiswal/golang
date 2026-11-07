package main

import "fmt"

// for loop is the only loop in golang
func main() {
	//  classic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 6
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	//  range loop
	for k := range 4 {
		if k == 2 {
			continue
		}
		fmt.Println(k)
	}

}

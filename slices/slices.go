package main

import (
	"fmt"
	"slices"
)

// main demonstrates common operations on slices in Go.
func main() {
	// A nil slice has a length of zero and no backing array.
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	// make creates a slice with a length of 3 and a capacity of 5.
	var num = make([]int, 3, 5)
	fmt.Println(num)
	fmt.Println(num == nil)
	fmt.Println(cap(num)) // capacity is the maximum number of elements before reallocation.

	num = append(num, 1)
	fmt.Println(num)
	fmt.Println(cap(num))
	num = append(num, 2)
	fmt.Println(num)
	fmt.Println(cap(num))
	num = append(num, 3) // at this point, slice capacity will be doubled to accommodate the new elements.
	fmt.Println(num)
	fmt.Println(cap(num))

	data := []int{1, 2, 3, 4}
	fmt.Println(data)

	// Append elements to a slice and then copy them into another slice.
	var slice1 = make([]int, 2, 3)
	slice1 = append(slice1, 1, 2, 3, 4)
	var slice2 = make([]int, len(slice1))

	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)

	// Slice operators select a sub-slice from the original slice.
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3[1:4])

	// slices.Equal compares two slices element by element.
	var slice4 = []int{1, 2, 3}
	var slice5 = []int{1, 2, 4}
	fmt.Println(slices.Equal(slice4, slice5))
}

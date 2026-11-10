package main 

import "fmt"

func main() {
	var arr [5]int

	// array length
	fmt.Println("Size of array is :", len(arr));

	arr[0] = 10
	fmt.Println("Element at index 0 is :", arr[0]);
	fmt.Println(arr); // zeroed values of array

	var values [4]bool

	values[0] = true
	// values[1] = false
	fmt.Println("Values of array is :", values);

	nums := [3]int{1, 2, 3}
	fmt.Println("Values of nums array is :", nums);

	data := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("Values of data array is :", data);

	// fixed size, that is predictable
	// memory optimization
	// constant time access
}
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	sum := 0

	for i, num := range nums {
		fmt.Println("Index:", i, "Value:", num)
		sum += num
	}
	fmt.Println(sum)

	names := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}

	for key, value := range names {
		fmt.Println("Key:", key, "Value:", value)
	}

	for i, c := range "golang" {
		fmt.Println("Index:", i, "Value:", string(c))
	}
}

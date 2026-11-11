package main

import (
	"fmt"
	"maps"
)

func main() {
	// simple map
	mapData := make(map[string]string)

	mapData["name"] = "John"
	mapData["age"] = "30"

	fmt.Println(mapData["name"], mapData["age"])
	fmt.Println(mapData["info"]) // This will print an empty string since "info" key does not exist
	// if key does not exist, it will return the zero value (like empty string for string type, 0 for int type, false for bool type, etc.)

	data := map[string]int{
		"phone": 123456789,
		"age":   30,
	}
	fmt.Println(data["phone"], data["age"], len(data))
	delete(data, "age")                       // delete key "age" from map
	fmt.Println(data, data["age"], len(data)) // "age" key will return 0 since it has been deleted
	clear(data)
	fmt.Println(data, len(data)) // map will be empty after clear

	values := map[string]int{
		"phone": 3,
		"cards": 5,
	}

	v, ok := values["phone"]
	fmt.Println(v, ok)

	if ok {
		fmt.Println("Ok")
	} else {
		fmt.Println("Not Ok")
	}

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 3}

	fmt.Println(maps.Equal(m1, m2)) // false, because the values for key "b" are different
}

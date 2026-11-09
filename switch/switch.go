package main

import "fmt"

func main() {
	// basic switch statement
	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	// multiple cases in a single case statement
	grade := "B"

	switch grade {
	case "A", "B":
		fmt.Println("Good job!")
	case "C", "D":
		fmt.Println("You can do better.")
	default:
		fmt.Println("Invalid grade")
	}

	// switch with no expression (acts like if-else)
	score := 85

	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 75:
		fmt.Println("Good")
	case score >= 50:
		fmt.Println("Average")
	default:
		fmt.Println("Fail")
	}

	typeIdentifier := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer:", t)
		case string:
			fmt.Println("String: ", t)
		default:
			fmt.Println("Unknown type")
		}
	}

	typeIdentifier("Raju") // Output: String: Raju
	typeIdentifier(42)     // Output: Integer: 42
	typeIdentifier(3.14)   // Output: Unknown type
}

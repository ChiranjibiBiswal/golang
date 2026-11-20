package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println("Printing number:", i)
		}(i)
	}

	time.Sleep(time.Second * 1) // time.Sleep is a temporary way to keep main alive

}

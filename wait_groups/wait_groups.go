package main

import (
	"fmt"
	"sync"
)

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Running task:", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}

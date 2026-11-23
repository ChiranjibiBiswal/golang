package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increment(w *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		w.Done()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go myPost.increment(&wg)
	}
	wg.Wait()
	fmt.Println("Post after increment:", myPost.views)
}
